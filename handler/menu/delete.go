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

func Delete(c *gin.Context) {
	var r DeleteRequest
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
	errNo = menuService.Delete(menuModel)
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	for _, roleModel := range menuModel.Roles {
		_ = inject.Obj.Common.RoleAPI.LoadPolicy(roleModel.ID)
	}
	handler.SendResponse(c, nil, nil)
}
