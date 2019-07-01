package menu

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/menuservice"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	menuService := menuservice.Menu{
		ID: r.ID,
	}
	menu, errNo := menuService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		ID:         menu.ID,
		ParentID:   menu.ParentID,
		ParentName: "",
		Name:       menu.Name,
		Url:        menu.Url,
		Perms:      menu.Perms,
		Type:       menu.Type,
		Icon:       menu.Icon,
		OrderNum:   menu.OrderNum,
		Open:       0,
		CreateTime: menu.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

func GetMenuNav(c *gin.Context) {
	userId, _ := c.Get("userID")
	menuService := menuservice.Menu{}
	list, permissions, errNo := menuService.GetMenuNavByUserId(userId.(uint64))
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, map[string]interface{}{
		"menuList":    list,
		"permissions": permissions,
	})
}
