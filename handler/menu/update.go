package menu

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/menuservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
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
	menuModel.ParentID = r.ParentID
	menuModel.Name = r.Name
	menuModel.URL = r.URL
	menuModel.Perms = r.Perms
	menuModel.Type = r.Type
	menuModel.Icon = r.Icon
	menuModel.OrderNum = r.OrderNum
	menuModel.IsTab = r.IsTab
	menuModel.Status = r.Status
	errNo = menuService.Edit(menuModel)
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	for _, roleModel := range menuModel.Roles {
		_ = inject.Obj.Common.RoleAPI.LoadPolicy(roleModel.ID)
	}
	handler.SendResponse(c, nil, nil)
}
