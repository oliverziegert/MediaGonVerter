package service

import (
	"encoding/base64"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"pc-ziegert.de/media_service/common/config"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	cs "pc-ziegert.de/media_service/service/client/core-service"
	"pc-ziegert.de/media_service/service/db/repo"
	"pc-ziegert.de/media_service/service/mq"
	"pc-ziegert.de/media_service/service/utils"
	"strings"
	"time"
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

func (i *ImageService) GetEncryptedToken(ctx *gin.Context, token *string) (*string, *e.Error) {
	et, err := splitToken(*token)
	if err != nil {
		return nil, err
	}

	return &et[1], nil
}

func (i *ImageService) GetTenantDomain(ctx *gin.Context, token *string) (*string, *e.Error) {
	tokenParts, err := splitToken(*token)
	if err != nil {
		return nil, err
	}
	tenantDomain, error := base64.StdEncoding.DecodeString(tokenParts[0])
	if error != nil {
		err := e.WrapError(e.ValIdInvalid, "Invalid token variable. (Token can't be base64 decoded.)", error)
		log.Debug(err.StackTrace())
		return nil, err
	}
	if len(tenantDomain) <= 1 {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (tenantDomain can't be less then 1 char.)")
		log.Debug(err.StackTrace())
		return nil, err
	}
	if !govalidator.IsDNSName(string(tenantDomain)) {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (tenantDomain is not a valid DNS name.)")
		log.Debug(err.StackTrace())
		panic(err)
	}

	domain := string(tenantDomain)
	return &domain, nil
}

func (i *ImageService) SplitSize(ctx *gin.Context, size *string) (uint64, uint64, *e.Error) {
	sizeParts := strings.Split(*size, "x")
	if len(sizeParts) != 2 {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. (Size does not contain two parts.)")
		log.Debug(err.StackTrace())
		return 0, 0, err
	}

	height := utils.StringToUint64(sizeParts[1])
	width := utils.StringToUint64(sizeParts[0])

	checkDimension(height)
	checkDimension(width)

	return width, height, nil
}

func (i *ImageService) DecryptToken(ctx *gin.Context) (*m.MediaToken, *e.Error) {
	encryptedToken := ctx.GetString(constant.ContextKeyEncryptedToken)
	if encryptedToken == "" {
		err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
		log.Debug(err.StackTrace())
		return nil, err
	}

	csc := cs.NewCoreSerivceClient(ctx, i.conf)
	mdto := cs.MediaDecryptTokenOpts{XSdsServiceToken: i.conf.Service.Internal.Communication.Auth.Token}
	mediaToken, err := csc.InternalApi.MediaDecryptToken(ctx, encryptedToken, &mdto)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return nil, err
	}

	mt := tokenToMediaToken(mediaToken, encryptedToken, ctx.GetString(constant.ContextKeyDracoonForwardedHostHeader))

	return mt, nil

}

func (i *ImageService) GetImageByToken(ctx *gin.Context, mediaToken *m.MediaToken, width uint64, height uint64, crop bool) (*m.Image, *e.Error) {
	ex := constant.ImageInitialExpiration
	image, err := mediaTokenToImage(ctx, mediaToken, width, height, crop)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Could not image from media token", err)
		log.Debug(err.StackTrace())
		return nil, err
	}

	exists, err := i.iRepo.ExistsImage(ctx, image)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Could not check if image is cached", err)
		log.Debug(err.StackTrace())
		return nil, err
	}

	if exists {
		err = i.iRepo.GetImage(ctx, image)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Could not check if image is cached", err)
			log.Debug(err.StackTrace())
			return nil, err
		}

		c, _ := image.GetConversionBySize(width, height, crop)
		if c != nil {
			if c.Expires.Unix() > (time.Now().Unix() + 60) {
				return image, nil
			}
			c.State = m.ConversionStateUnknown
			ex = i.conf.Data.ExpirationDuration
		} else {
			image.Conversions = append(image.Conversions, m.NewConversion(width, height, crop))
		}
	}

	err = i.iRepo.SetImage(ctx, image, &ex)
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

	err = i.iRepo.UpdateImage(ctx, image)
	if err != nil {
		return nil, err
	}

	err = i.downloadAndConvert(ctx, image)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (i *ImageService) getS3DownloadUrl(ctx *gin.Context, image *m.Image) *e.Error {
	// Get jwt token claims
	cc, exists := GetJwtClaimsFromContext(ctx)
	if !exists {
		err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
		log.Debug(err.StackTrace())
		return err
	}

	csc := cs.NewCoreSerivceClient(ctx, i.conf)
	ffido := cs.FilesFileIdDownloadsOpts{
		XDcInternalServiceToken: i.conf.Service.Internal.Communication.Auth.Token,
		UserId:                  cc.UserId,
	}

	filesFileIDDownloads, err := csc.InternalApi.FilesFileIDDownloads(ctx, image.NodeId, &ffido)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}

	image.S3DownloadUrl = filesFileIDDownloads.DownloadUrl
	return nil
}

func (i *ImageService) downloadAndConvert(ctx *gin.Context, image *m.Image) *e.Error {
	pub, err := i.mq.NewPublishing(image)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	routingKey, err := utils.GetRabbitMQWorkerRoutingKeyFromNodeType(image.NodeType)
	if err != nil {
		return err
	}
	err = i.mq.Publish(
		constant.RabbitMQExchangeName,
		routingKey,
		false,
		false,
		*pub)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

func (i *ImageService) generateS3UploadUrls(ctx *gin.Context, image *m.Image) *e.Error {
	var imgUpdate = false
	cc, exists := GetJwtClaimsFromContext(ctx)
	if !exists {
		err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
		log.Debug(err.StackTrace())
		return err
	}

	for _, conversion := range image.Conversions {
		if conversion.State != m.ConversionStateCached {
			key := fmt.Sprintf(constant.S3KeyTemplate,
				image.NodeId%10,
				cc.TenantUuid,
				cc.CustomerUuid,
				image.NodeId,
				conversion.Width,
				conversion.Height,
				conversion.Crop)
			req, expires, err := utils.GenerateS3PresignUploadUrl(ctx, i.conf, key, image.NodeType)

			if err != nil {
				err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
				log.Debug(err.StackTrace())
				return err
			}
			conversion.S3UploadUrl = req
			conversion.Expires = expires
			imgUpdate = true
		}
	}
	if imgUpdate {
		err := i.iRepo.UpdateImage(ctx, image)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			log.Debug(err.StackTrace())
			return err
		}
	}
	return nil
}

func (i *ImageService) GenerateS3DownloadUrl(ctx *gin.Context, image *m.Image, conversion *m.Conversion) (string, *e.Error) {
	// Get jwt token claims
	cc, exists := GetJwtClaimsFromContext(ctx)
	if !exists {
		err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
		log.Debug(err.StackTrace())
		return "", err
	}

	key := fmt.Sprintf(constant.S3KeyTemplate,
		image.NodeId%10,
		cc.TenantUuid,
		cc.CustomerUuid,
		image.NodeId,
		conversion.Width,
		conversion.Height,
		conversion.Crop)
	req, err := utils.GenerateS3PresignDownloadUrl(ctx, i.conf, key, image.NodeType)

	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return "", err
	}
	return req.URL, nil
}

func (i *ImageService) checkConversionExists(ctx *gin.Context, image *m.Image, width uint64, height uint64, crop bool) bool {
	for _, conversion := range image.Conversions {
		if conversion.Crop == crop &&
			conversion.Width == width &&
			conversion.Height == height {
			return true
		}
	}
	return false
}

func splitToken(token string) ([]string, *e.Error) {
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 2 {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Token does not contain two parts.)")
		log.Debug(err.StackTrace())
		return nil, err
	}
	return tokenParts, nil
}

func checkDimension(i uint64) bool {
	if i >= 1 {
		return true
	}
	return false
}

func tokenToMediaToken(mediaToken *m.MediaDecryptedToken, token string, tenantDomain string) *m.MediaToken {
	tokenParts := strings.Split(mediaToken.Token, "|")
	return m.NewMediaToken(
		utils.StringToUint64(tokenParts[0]),
		tokenParts[1],
		tokenParts[4],
		tenantDomain,
		utils.StringToUint64(tokenParts[2]),
		token,
		tokenParts[3],
	)
}

func mediaTokenToImage(ctx *gin.Context, mediaToken *m.MediaToken, width uint64, height uint64, crop bool) (*m.Image, *e.Error) {
	h := utils.GetHostname()
	cc, exists := GetJwtClaimsFromContext(ctx)
	if !exists {
		err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
		log.Debug(err.StackTrace())
		return nil, err
	}

	image := m.NewImage(
		cc.TenantUuid,
		mediaToken.NodeId)

	image.ServiceName = *h
	image.NodeType = mediaToken.NodeType
	image.NodeSize = mediaToken.NodeSize
	image.Conversions = []*m.Conversion{
		m.NewConversion(width, height, crop),
	}
	return image, nil
}
