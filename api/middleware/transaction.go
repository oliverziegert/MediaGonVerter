package middleware

import (
	"github.com/gin-gonic/gin"
	"pc-ziegert.de/media_service/constant"
)

// TransactionMiddleware creates the security context.
type TransactionMiddleware struct {
}

// NewTransactionMiddleware  create a new TransactionMiddleware
func NewTransactionMiddleware() *TransactionMiddleware {
	return &TransactionMiddleware{}
}

// ServeHTTP processes requests.
func (m *TransactionMiddleware) ServeHTTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(constant.ContextKeySessionHolder, "")
	}
}
