package middleware

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/service/service"
	"strconv"
	"strings"
)

// SecurityMiddleware creates the security context.
type SecurityMiddleware struct{}

// NewSecurityMiddleware create a new HeaderMiddleware.
func NewSecurityMiddleware() *SecurityMiddleware {
	return &SecurityMiddleware{}
}

// ServeHTTP processes requests.
func (m *SecurityMiddleware) ServeHTTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		forwardedProtoHeader := ctx.GetHeader(constant.DracoonForwardedProtoHeader)
		if forwardedProtoHeader == "" {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("%v header not provided", constant.DracoonForwardedProtoHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		if !(strings.EqualFold(forwardedProtoHeader, "https") || strings.EqualFold(forwardedProtoHeader, "http")) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Header %v should be http or https, but is: %s", constant.DracoonForwardedProtoHeader, forwardedProtoHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		ctx.Set(constant.ContextKeyDracoonForwardedProtoHeader, forwardedProtoHeader)

		forwardedHostHeader := ctx.GetHeader(constant.DracoonForwardedHostHeader)
		if forwardedHostHeader == "" {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("%v header not provided", constant.DracoonForwardedHostHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		if !govalidator.IsDNSName(forwardedHostHeader) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Invalid header %v, is not a valid DNS name.", constant.DracoonForwardedHostHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		ctx.Set(constant.ContextKeyDracoonForwardedHostHeader, forwardedHostHeader)

		forwardedPortHeader := ctx.GetHeader(constant.DracoonForwardedPortHeader)
		if forwardedPortHeader == "" {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("%v header not provided", constant.DracoonForwardedPortHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		if !govalidator.IsInt(forwardedPortHeader) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Invalid header %v, is not int.", constant.DracoonForwardedPortHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		forwardedPortHeaderInt, err := strconv.Atoi(forwardedPortHeader)
		if err != nil && (forwardedPortHeaderInt < 1 || forwardedPortHeaderInt > 65535) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Invalid header %v, is not between 1 and 65535.", constant.DracoonForwardedPortHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		ctx.Set(constant.ContextKeyDracoonForwardedPortHeader, uint64(forwardedPortHeaderInt))

		clientAddressHeader := ctx.GetHeader(constant.DracoonClientAddressHeader)
		if clientAddressHeader == "" {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("%v header not provided", constant.DracoonClientAddressHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		if !(govalidator.IsIPv4(clientAddressHeader) || govalidator.IsIPv6(clientAddressHeader)) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Invalid header %v, is not a IPv4 or IPv6 address.", constant.DracoonClientAddressHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		ctx.Set(constant.ContextKeyDracoonClientAddressHeader, clientAddressHeader)

		clientUserAgentHeader := ctx.GetHeader(constant.DracoonClientUserAgentHeader)
		if clientUserAgentHeader == "" {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("%v header not provided", constant.DracoonClientUserAgentHeader))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		ctx.Set(constant.ContextKeyDracoonClientUserAgentHeader, clientUserAgentHeader)

		ctx.Next()
	}
}

// AuthCheckMiddleware ensures that a user was authenticated.
type AuthCheckMiddleware struct {
	jwts *service.JWTService
}

// NewAuthCheckMiddleware create a new AuthCheckMiddleware.
func NewAuthCheckMiddleware(jwts *service.JWTService) *AuthCheckMiddleware {
	return &AuthCheckMiddleware{
		jwts: jwts,
	}
}

// ServeHTTP processes requests.
func (m *AuthCheckMiddleware) ServeHTTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := strings.Trim(authHeader[len(BearerSchema):], " ")

		token, mapClaims, err := m.jwts.ValidateToken(tokenString)
		if !token.Valid {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			log.Debug(err.StackTrace())
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		if !mapClaims.VerifyIssuer(constant.JWTIssuer, true) {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			log.Debug(err.StackTrace())
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		cc, err := m.jwts.Claims(&mapClaims)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			log.Debug(err.StackTrace())
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Set(constant.ContextKeyJWT, token)
		ctx.Set(constant.ContextKeyJWTTokenClaims, cc)

		ctx.Next()
	}
}
