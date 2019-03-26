package middleware

import (
	"github.com/gin-gonic/gin"
	"cladmin/pkg/token"
	. "cladmin/handler"
	"cladmin/pkg/errno"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			SendResponseUnauthorized(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
