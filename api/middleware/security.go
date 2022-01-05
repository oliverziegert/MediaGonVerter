package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecurityMiddleware creates the security context.
type SecurityMiddleware struct{}

// NewSecurityMiddleware create a new SecurityMiddleware.
func NewSecurityMiddleware() *SecurityMiddleware {
	return &SecurityMiddleware{}
}

// ServeHTTP processes requests.
func (m *SecurityMiddleware) ServeHTTP() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

// AuthCheckMiddleware ensures that a user was authenticated.
type AuthCheckMiddleware struct {
}

// NewAuthCheckMiddleware create a new AuthCheckMiddleware.
func NewAuthCheckMiddleware() *AuthCheckMiddleware {
	return &AuthCheckMiddleware{}
}

// ServeHTTP processes requests.
func (m *AuthCheckMiddleware) ServeHTTP() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
