package permission

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"github.com/chenglongcl/log"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")
		superAdmin := c.GetBool("superAdmin")
		if b, err := inject.Obj.Enforcer.EnforceSafe(username, c.Request.URL.Path, superAdmin); err != nil {
			log.Fatal("Casbin EnforceSafe Error", err)
			handler.SendResponseUnauthorized(c, errno.ErrCasbin, nil)
			c.Abort()
			return
		} else if !b {
			handler.SendResponseForbidden(c, errno.ErrNotPermission, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
