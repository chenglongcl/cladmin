package menu

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/menuservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	menuService := menuservice.Menu{
		ParentID: r.ParentID,
		Name:     r.Name,
		Url:      r.Url,
		Perms:    r.Perms,
		Type:     r.Type,
		Icon:     r.Icon,
		OrderNum: r.OrderNum,
	}
	if errNo := menuService.Add(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
