package service

import (
	"encoding/base64"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"pc-ziegert.de/media_service/config"
	"pc-ziegert.de/media_service/constant"
	e "pc-ziegert.de/media_service/error"
	"pc-ziegert.de/media_service/log"
	"pc-ziegert.de/media_service/model"
	"strconv"
	"strings"
)

// ImageService contains image related logic.
type ImageService struct {
	conf *config.Config
	//	iRepo *repo.ImageRepo
}

// NewImageService NewUserService create a new user service.
func NewImageService(conf *config.Config) *ImageService {
	return &ImageService{conf: conf}
}

func (i *ImageService) GetEncryptedToken(ctx *gin.Context, token string) string {
	ctx.Set("encryptedToken", splitToken(token)[1])
	return splitToken(token)[1]
}

func (i *ImageService) GetTenantDomain(ctx *gin.Context, token string) string {
	tokenParts := splitToken(token)
	tenantDomain, err := base64.StdEncoding.DecodeString(tokenParts[0])
	if err != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Token can't be base64 decoded.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	if len(tenantDomain) <= 1 {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (tenantDomain can't be less then 1 char.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	if !govalidator.IsDNSName(string(tenantDomain)) {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (tenantDomain is not a valid DNS name.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	ctx.Set("tenantDomain", string(tenantDomain))
	return string(tenantDomain)
}

func (i *ImageService) SplitSize(ctx *gin.Context, size string) (int, int) {
	sizeParts := strings.Split(size, "x")
	if len(sizeParts) != 2 {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. (Size does not contain two parts.)")
		log.Debug(err.StackTrace())
		panic(err)
	}

	height := stringToInt(sizeParts[1])
	width := stringToInt(sizeParts[0])

	checkDimension(height)
	checkDimension(width)

	return width, height
}

func (i *ImageService) DecryptToken(ctx *gin.Context, token string) *model.MediaToken {
	protocol := i.conf.Service.Core.Protocol
	host := i.conf.Service.Core.Host
	port := i.conf.Service.Core.Port
	apiKey := i.conf.Service.Internal.Communication.Auth.Token
	url := fmt.Sprintf(constant.CoreServiceMediaTokenDecryptTemplate, protocol, host, port, token)

	c := NewClient(url, apiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil
	}

	res, err := c.sendRequest(ctx, req)

	resParts := strings.Split(res.Data, "|")

	if len(resParts) != 5 {
		return nil
	}

	return model.NewMediaToken(
		stringToUint64(resParts[0]),
		resParts[1],
		resParts[4],
		ctx.GetString("tenantDomain"),
		stringToUint64(resParts[2]),
		token,
		resParts[3],
	)
}

func splitToken(token string) []string {
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 2 {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Token does not contain two parts.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return tokenParts
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. ("+s+" cant be converted to int.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return i
}

func stringToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. ("+s+" cant be converted to int.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return i
}

func checkDimension(i int) bool {
	ok := i >= 1
	if !ok {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. ("+string(i)+" is not greater or equal to 1)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return ok
}
