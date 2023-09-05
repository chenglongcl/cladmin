package bulletin

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/bulletinservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	bulletinService := bulletinservice.NewBulletinService(c)
	bulletinService.Title = r.Title
	bulletinService.Tag = r.Tag
	bulletinService.Content = r.Content
	if _, errNo := bulletinService.Add(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
