package middleware

import (
	"encoding/base64"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"pc-ziegert.de/media_service/common/constant"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/service/service"
	u "pc-ziegert.de/media_service/service/utils"
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
			l.Debug(err.StackTrace())
			panic(err)
		}
		if !(strings.EqualFold(forwardedProtoHeader, "https") || strings.EqualFold(forwardedProtoHeader, "http")) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Header %v should be http or https, but is: %s", constant.DracoonForwardedProtoHeader, forwardedProtoHeader))
			l.Debug(err.StackTrace())
			panic(err)
		}
		ctx.Set(constant.ContextKeyDracoonForwardedProtoHeader, forwardedProtoHeader)

		forwardedHostHeader := ctx.GetHeader(constant.DracoonForwardedHostHeader)
		if forwardedHostHeader == "" {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("%v header not provided", constant.DracoonForwardedHostHeader))
			l.Debug(err.StackTrace())
			panic(err)
		}
		if !govalidator.IsDNSName(forwardedHostHeader) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Invalid header %v, is not a valid DNS name.", constant.DracoonForwardedHostHeader))
			l.Debug(err.StackTrace())
			panic(err)
		}
		ctx.Set(constant.ContextKeyDracoonForwardedHostHeader, forwardedHostHeader)

		forwardedPortHeader := ctx.GetHeader(constant.DracoonForwardedPortHeader)
		if forwardedPortHeader == "" {
			if forwardedProtoHeader == "https" {
				forwardedPortHeader = "443"
			}
			if forwardedProtoHeader == "http" {
				forwardedPortHeader = "80"
			}
		}
		if !govalidator.IsInt(forwardedPortHeader) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Invalid header %v, is not int.", constant.DracoonForwardedPortHeader))
			l.Debug(err.StackTrace())
			panic(err)
		}
		forwardedPortHeaderInt, err := strconv.Atoi(forwardedPortHeader)
		if err != nil && (forwardedPortHeaderInt < 1 || forwardedPortHeaderInt > 65535) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Invalid header %v, is not between 1 and 65535.", constant.DracoonForwardedPortHeader))
			l.Debug(err.StackTrace())
			panic(err)
		}
		ctx.Set(constant.ContextKeyDracoonForwardedPortHeader, uint64(forwardedPortHeaderInt))

		clientAddressHeader := ctx.GetHeader(constant.DracoonClientAddressHeader)
		if clientAddressHeader == "" {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("%v header not provided", constant.DracoonClientAddressHeader))
			l.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		if !(govalidator.IsIPv4(clientAddressHeader) || govalidator.IsIPv6(clientAddressHeader)) {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("Invalid header %v, is not a IPv4 or IPv6 address.", constant.DracoonClientAddressHeader))
			l.Debug(err.StackTrace())
			panic(err)
		}
		ctx.Set(constant.ContextKeyDracoonClientAddressHeader, clientAddressHeader)

		clientUserAgentHeader := ctx.GetHeader(constant.DracoonClientUserAgentHeader)
		if clientUserAgentHeader == "" {
			err := e.NewError(e.ValIdInvalid, fmt.Sprintf("%v header not provided", constant.DracoonClientUserAgentHeader))
			l.Debug(err.StackTrace())
			panic(err)
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
			l.Debug(err.StackTrace())
			panic(err)
		}
		if !mapClaims.VerifyIssuer(constant.JWTIssuer, true) {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			l.Debug(err.StackTrace())
			panic(err)
		}

		cc, err := m.jwts.Claims(&mapClaims)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			l.Debug(err.StackTrace())
			panic(err)
		}

		ctx.Set(constant.ContextKeyJWT, token)
		ctx.Set(constant.ContextKeyJWTTokenClaims, cc)

		ctx.Next()
	}
}

// TokenCheckMiddleware ensures that a user was authenticated.
type TokenCheckMiddleware struct{}

// NewTokenCheckMiddleware create a new TokenCheckMiddleware.
func NewTokenCheckMiddleware() *TokenCheckMiddleware {
	return &TokenCheckMiddleware{}
}

// ServeHTTP processes requests.
func (m *TokenCheckMiddleware) ServeHTTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, err := u.GetTokenPathVar(ctx)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			l.Debug(err.StackTrace())
			panic(err)
		}

		tp, err := u.SplitMediaToken(t)
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			l.Debug(err.StackTrace())
			panic(err)
		}

		td, err := u.GetTenantDomainFromMediaToken(tp[0])
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			l.Debug(err.StackTrace())
			panic(err)
		}

		stp, error := base64.StdEncoding.DecodeString(tp[1])
		if error != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", error)
			l.Debug(err.StackTrace())
			panic(err)
		}

		stpp := strings.Split(string(stp), "|")
		if len(stpp) != 3 {
			err := e.NewError(e.ValIdInvalid, "Invalid token variable. (Token does not contain two parts.)")
			l.Debug(err.StackTrace())
			panic(err)
		}

		cc := &service.CustomClaims{
			TenantUuid:   *td,
			CustomerUuid: "token",
			UserId:       u.StringToUint64(stpp[0]),
		}
		ctx.Set(constant.ContextKeyJWTTokenClaims, cc)

		ctx.Next()
	}
}
