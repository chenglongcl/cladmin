package menu

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/menuservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	menuService := menuservice.NewMenuService(c)
	menuModel, errNo := menuService.Get([]field.Expr{
		cladminquery.Q.SysMenu.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysMenu.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if menuModel == nil || menuModel.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, GetResponse{
		ID:         menuModel.ID,
		ParentID:   menuModel.ParentID,
		ParentName: "",
		Name:       menuModel.Name,
		Url:        menuModel.URL,
		Perms:      menuModel.Perms,
		Type:       menuModel.Type,
		Icon:       menuModel.Icon,
		OrderNum:   menuModel.OrderNum,
		IsTab:      menuModel.IsTab,
		Status:     menuModel.Status,
		CreateTime: menuModel.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

func GetMenuNav(c *gin.Context) {
	userID := c.GetUint64("userID")
	menuService := menuservice.NewMenuService(c)
	menus, errNo := menuService.GetMenuNavByUserID(userID)
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, menus)
}

func GetPermissions(c *gin.Context) {
	userID := c.GetUint64("userID")
	menuService := menuservice.NewMenuService(c)
	permissions, errNo := menuService.GetPermissionsByUserID(userID)
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, permissions)
}
