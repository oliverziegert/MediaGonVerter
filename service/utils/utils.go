package utils

import (
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"regexp"
	"strconv"
	"strings"
)

var (
	hostname = ""
	letters  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. ("+s+" cant be converted to int.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return i
}

func StringToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. ("+s+" cant be converted to uint64.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return i
}

func StringToBool(s string) bool {
	o, err := strconv.ParseBool(s)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid string variable. ("+s+" cant be converted to bool.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return o
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetHostname() *string {
	if hostname != "" {
		return &hostname
	}

	h, err := os.Hostname()
	if err != nil {
		hostname = randSeq(10)
	} else {
		hostname = h
	}

	return &hostname
}

func GetTokenPathVar(ctx *gin.Context) (*string, *e.Error) {
	token := ctx.Param("token")
	if token == "" {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Varible missing.)")
		log.Debug(err.StackTrace())
		return nil, err
	}
	return &token, nil
}

func GetSizePathVar(ctx *gin.Context) (*string, *e.Error) {
	size := ctx.Param("size")
	if size == "" {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Varible missing.)")
		log.Debug(err.StackTrace())
		return nil, err
	}
	match, _ := regexp.MatchString("^\\d{1,5}x\\d{1,5}$", size)
	if !match {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. (Varible does not match \\d{1,5}x\\d{1,5}.)")
		log.Debug(err.StackTrace())
		return nil, err
	}
	return &size, nil
}

func GetCropVar(ctx *gin.Context) bool {
	value := ctx.DefaultQuery("crop", "false")
	switch value {
	case "true":
		return true
	case "false":
		return false
	}
	return false
}

func SplitMediaToken(token *string) ([]string, *e.Error) {
	tokenParts := strings.Split(*token, ".")
	if len(tokenParts) != 2 {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Token does not contain two parts.)")
		log.Debug(err.StackTrace())
		return nil, err
	}
	return tokenParts, nil
}

func GetTenantDomainFromMediaToken(token string) (*string, *e.Error) {
	tenantDomain, error := base64.StdEncoding.DecodeString(token)
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
