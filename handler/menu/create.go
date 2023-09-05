package menu

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/menuservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	menuService := menuservice.NewMenuService(c)
	menuService.ParentID = r.ParentID
	menuService.Name = r.Name
	menuService.URL = r.URL
	menuService.Perms = r.Perms
	menuService.Type = r.Type
	menuService.Icon = r.Icon
	menuService.OrderNum = r.OrderNum
	if _, errNo := menuService.Add(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
