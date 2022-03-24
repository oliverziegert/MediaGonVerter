package service

import (
	"github.com/gin-gonic/gin"
	"pc-ziegert.de/media_service/common/constant"
)

func GetJwtClaimsFromContext(ctx *gin.Context) (*CustomClaims, bool) {
	// Get jwt token claims
	value, exists := ctx.Get(constant.ContextKeyJWTTokenClaims)
	if !exists {
		return nil, exists
	}

	// create tenant object from jwt token claim uuid
	return value.(*CustomClaims), exists
}
