package controller

import (
	"github.com/gin-gonic/gin"
	e "pc-ziegert.de/media_service/error"
	"pc-ziegert.de/media_service/log"
	"regexp"
)

type Error struct {
	string error
}

func getTokenPathVar(c *gin.Context) string {
	token := c.Param("token")
	if token == "" {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Varible missing.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return token
}

func getSizePathVar(c *gin.Context) string {
	size := c.Param("size")
	if size == "" {
		err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Varible missing.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	match, _ := regexp.MatchString("^\\d{1,5}x\\d{1,5}$", size)
	if !match {
		err := e.NewError(e.ValIdInvalid, "Invalid size variable. (Varible does not match \\d{1,5}x\\d{1,5}.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return size
}

func getCropVar(context *gin.Context) bool {
	value := context.DefaultQuery("crop", "false")
	switch value {
	case "true":
		return true
	case "false":
		return false
	}
	return false
}
