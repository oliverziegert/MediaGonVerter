package service

import (
	"encoding/base64"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	error2 "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	model2 "pc-ziegert.de/media_service/common/model"
	cs "pc-ziegert.de/media_service/service/client/core-service"
	"pc-ziegert.de/media_service/service/db/repo"
	"pc-ziegert.de/media_service/service/mq"
	"pc-ziegert.de/media_service/service/utils"
	"strings"
)

// ImageService contains image related logic.
type ImageService struct {
	conf  *config.Config
	iRepo *repo.ImageRepo
	mq    *mq.RabbitMQ
}

// NewImageService NewUserService create a new user service.
func NewImageService(conf *config.Config, iRepo *repo.ImageRepo, mq *mq.RabbitMQ) *ImageService {
	return &ImageService{
		conf:  conf,
		iRepo: iRepo,
		mq:    mq,
	}
}

func (i *ImageService) GetEncryptedToken(ctx *gin.Context, token string) string {
	return splitToken(token)[1]
}

func (i *ImageService) GetTenantDomain(ctx *gin.Context, token string) string {
	tokenParts := splitToken(token)
	tenantDomain, err := base64.StdEncoding.DecodeString(tokenParts[0])
	if err != nil {
		err := error2.NewError(error2.ValIdInvalid, "Invalid token variable. (Token can't be base64 decoded.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	if len(tenantDomain) <= 1 {
		err := error2.NewError(error2.ValIdInvalid, "Invalid token variable. (tenantDomain can't be less then 1 char.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	if !govalidator.IsDNSName(string(tenantDomain)) {
		err := error2.NewError(error2.ValIdInvalid, "Invalid token variable. (tenantDomain is not a valid DNS name.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return string(tenantDomain)
}

func (i *ImageService) SplitSize(ctx *gin.Context, size string) (uint64, uint64) {
	sizeParts := strings.Split(size, "x")
	if len(sizeParts) != 2 {
		err := error2.NewError(error2.ValIdInvalid, "Invalid size variable. (Size does not contain two parts.)")
		log.Debug(err.StackTrace())
		panic(err)
	}

	height := utils.StringToUint64(sizeParts[1])
	width := utils.StringToUint64(sizeParts[0])

	checkDimension(height)
	checkDimension(width)

	return width, height
}

func (i *ImageService) DecryptToken(ctx *gin.Context, token string) *model2.MediaToken {
	csc := cs.NewCoreSerivceClient(ctx, i.conf)

	// ToDo: Error handling
	mediaToken, _, _ := csc.InternalApi.MediaDecryptToken(ctx, token,
		&cs.MediaDecryptTokenOpts{XSdsServiceToken: i.conf.Service.Internal.Communication.Auth.Token})

	return tokenToMediaToken(&mediaToken, token, ctx.GetString(constant.ContextKeyDracoonForwardedHostHeader))

}

func (i *ImageService) GetImage(ctx *gin.Context, mediaToken *model2.MediaToken, width uint64, height uint64, crop bool) (*model2.Image, *error2.Error) {
	image := mediaTokenToImage(mediaToken, width, height, crop)

	exists, err := i.iRepo.ExistsImage(ctx, image)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "Could not check if image is cached", err)
		log.Debug(err.StackTrace())
		return nil, err
	}

	if exists {
		err = i.iRepo.GetImage(ctx, image)
		return image, err
	}

	err = i.iRepo.SetImage(ctx, image)
	if err != nil {
		return nil, err
	}

	imageCheck := *image
	err = i.iRepo.GetImage(ctx, &imageCheck)
	if err != nil {
		return nil, err
	}

	if !(image.ServiceName == imageCheck.ServiceName) {
		return image, nil
	}

	err = i.getS3DownloadUrl(ctx, image)
	if err != nil {
		return nil, err
	}

	err = i.generateS3UploadUrls(ctx, image)
	if err != nil {
		return nil, err
	}

	err = i.iRepo.SetImage(ctx, image)
	if err != nil {
		return nil, err
	}

	err = i.downloadAndConvert(ctx, image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (i *ImageService) getS3DownloadUrl(ctx *gin.Context, image *model2.Image) *error2.Error {
	var err error
	csc := cs.NewCoreSerivceClient(ctx, i.conf)
	filesFileIDDownloads, _, _ := csc.InternalApi.FilesFileIDDownloads(
		ctx,
		image.NodeId,
		&cs.FilesFileIDDownloadsOpts{XDcInternalServiceToken: i.conf.Service.Internal.Communication.Auth.Token})
	image.S3DownloadUrl, err = url.Parse(filesFileIDDownloads.DownloadUrl)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

func (i *ImageService) downloadAndConvert(ctx *gin.Context, image *model2.Image) *error2.Error {

	pub, err := i.mq.NewPublishing(image)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	err = i.mq.Publish(
		constant.RabbitMQExchangeName,
		constant.RabbitMQWorkerRoutingKey,
		false,
		false,
		*pub)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

func (i *ImageService) generateS3UploadUrls(ctx *gin.Context, image *model2.Image) *error2.Error {
	for _, conversion := range image.Conversions {
		var claims *CustomClaims
		if c, exists := ctx.Get(constant.ContextKeyJWTTokenClaims); exists {
			claims = c.(*CustomClaims)
		} else {
			err := error2.NewError(error2.ValIdInvalid, "Failed to register a consumer.")
			return err
		}
		key := fmt.Sprintf(constant.S3KeyTemplate,
			image.NodeId%10,
			claims.TenantUuid,
			claims.CustomerUuid,
			image.NodeId,
			conversion.Width,
			conversion.Height)
		req, err := utils.GenerateS3PresignUploadUrl(ctx, i.conf, key)

		if err != nil {
			err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", err)
			log.Debug(err.StackTrace())
			return err
		}
		url, error := url.Parse(req.URL)
		if error != nil {
			err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", error)
			log.Debug(err.StackTrace())
			return err
		}
		conversion.S3UploadUrl = url
	}
	return nil
}

func splitToken(token string) []string {
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 2 {
		err := error2.NewError(error2.ValIdInvalid, "Invalid token variable. (Token does not contain two parts.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return tokenParts
}

func checkDimension(i uint64) bool {
	ok := i >= 1
	if !ok {
		err := error2.NewError(error2.ValIdInvalid, "Invalid size variable. ("+string(i)+" is not greater or equal to 1)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return ok
}

// ToDo: Error handling
func tokenToMediaToken(mediaToken *model2.MediaDecryptedToken, token string, tenantDomain string) *model2.MediaToken {
	tokenParts := strings.Split(mediaToken.Token, "|")
	return model2.NewMediaToken(
		utils.StringToUint64(tokenParts[0]),
		tokenParts[1],
		tokenParts[4],
		tenantDomain,
		utils.StringToUint64(tokenParts[2]),
		token,
		tokenParts[3],
	)
}

func mediaTokenToImage(mediaToken *model2.MediaToken, width uint64, height uint64, crop bool) *model2.Image {
	hostname, err := os.Hostname()
	if err != nil {
		//ToDo: Error handling
	}
	image := model2.NewImage(
		mediaToken.TenantDomain,
		mediaToken.NodeId)

	image.ServiceName = hostname
	image.NodeType = mediaToken.NodeType
	image.NodeSize = mediaToken.NodeSize
	image.Conversions = []*model2.Conversion{{
		Width:       width,
		Height:      height,
		Crop:        crop,
		S3UploadUrl: nil,
	}}
	return image
}
