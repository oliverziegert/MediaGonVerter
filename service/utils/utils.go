package utils

import (
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
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
		l.Debug(err.StackTrace())
		panic(err)
	}
	return i
}

func StringToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. ("+s+" cant be converted to uint64.)")
		l.Debug(err.StackTrace())
		panic(err)
	}
	return i
}

func StringToBool(s string) bool {
	o, err := strconv.ParseBool(s)
	if err != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid string variable. ("+s+" cant be converted to bool.)")
		l.Debug(err.StackTrace())
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
		l.Debug(err.StackTrace())
		return nil, err
	}
	return &token, nil
}

func GetSizePathVar(ctx *gin.Context) (*string, *e.Error) {
	size := ctx.Param("size")
	if size == "" {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Varible missing.)")
		l.Debug(err.StackTrace())
		return nil, err
	}
	match, _ := regexp.MatchString("^\\d{1,5}x\\d{1,5}$", size)
	if !match {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. (Varible does not match \\d{1,5}x\\d{1,5}.)")
		l.Debug(err.StackTrace())
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

// GetMediaTokenParts return individual parts of the given token
//(HostName bas64).(UserId|nodeId|{nodeId+userTokenSha256}  base64)
func GetMediaTokenParts(token *string) (*string, *string, *uint64, *uint64, *e.Error) {
	var hostName string
	var encryptedToken string
	var userId uint64
	var nodeId uint64

	tokenParts := strings.Split(*token, ".")
	if len(tokenParts) != 2 {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Token does not contain two parts.)")
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}

	firstTokenPart := tokenParts[0]
	secondTokenPart := tokenParts[1]
	hostNameByteArray, error := base64.StdEncoding.DecodeString(firstTokenPart)
	if error != nil {
		err := e.WrapError(e.ValIdInvalid, "Invalid token variable. (Token can't be base64 decoded.)", error)
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}
	hostName = string(hostNameByteArray)

	if len(hostName) <= 1 {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (hostName can't be less then 1 char.)")
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}

	if !govalidator.IsDNSName(hostName) {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (hostName is not a valid DNS name.)")
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}

	secondTokenPartByteArray, error := base64.StdEncoding.DecodeString(secondTokenPart)
	if error != nil {
		err := e.WrapError(e.ValIdInvalid, "Invalid token variable. (Token can't be base64 decoded.)", error)
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}
	encryptedToken = string(secondTokenPartByteArray)

	if !strings.Contains(encryptedToken, "|") {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Second token part does not contain |)")
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}

	secondTokenPartParts := strings.Split(encryptedToken, "|")
	if len(secondTokenPartParts) != 3 {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Token does not contain three parts.)")
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}

	userId = StringToUint64(secondTokenPartParts[0])
	if r := recover(); r != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (userId can not be extracted.)")
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}

	nodeId = StringToUint64(secondTokenPartParts[1])
	if r := recover(); r != nil {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (nodeId can not be extracted.)")
		l.Debug(err.StackTrace())
		return nil, nil, nil, nil, err
	}

	return &hostName, &encryptedToken, &userId, &nodeId, nil
}
