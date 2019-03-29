package menu

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/menu_service"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	menuService := menu_service.Menu{
		Id: r.Id,
	}
	menu, errNo := menuService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		Id:         menu.Id,
		ParentId:   menu.ParentId,
		Name:       menu.Name,
		Url:        menu.Url,
		Perms:      menu.Perms,
		Type:       menu.Type,
		Icon:       menu.Icon,
		OrderNum:   menu.OrderNum,
		CreateTime: menu.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
