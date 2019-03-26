package user

import (
	"github.com/gin-gonic/gin"
	"cladmin/pkg/token"
	. "cladmin/handler"
	"cladmin/pkg/errno"
)

func Refresh(c *gin.Context) {
	if ctx, err, t, e, r := token.ParseRefreshRequest(c); err != nil {
		SendResponseUnauthorized(c, errno.ErrTokenInvalid, nil)
	} else {
		SendResponse(c, nil, CreateResponse{
			Username:         ctx.Username,
			Token:            t,
			ExpiredAt:        e,
			RefreshExpiredAt: r,
		})
	}
}
