package middleware

import (
	"github.com/gin-gonic/gin"
)

// ErrorMiddleware creates the security context.
type ErrorMiddleware struct{}

// NewErrorMiddleware  create a new ErrorMiddleware
func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

// ServeHTTP processes requests.
func (m *ErrorMiddleware) ServeHTTP() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
