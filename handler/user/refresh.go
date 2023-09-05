package user

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/pkg/token"
	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	if ctx, err, t, e, r := token.ParseRefreshRequest(c); err != nil {
		handler.SendResponseForbidden(c, errno.ErrTokenInvalid, nil)
	} else {
		handler.SendResponse(c, nil, CreateResponse{
			Username:         ctx.Username,
			Token:            t,
			ExpiredAt:        e,
			RefreshExpiredAt: r,
		})
	}
}
