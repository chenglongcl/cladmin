package menu

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/service"
	"cladmin/service/menuservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func List(c *gin.Context) {
	menuService := menuservice.NewMenuService(c)
	info, _, errNo := menuService.InfoList(&service.ListParams{
		PS: util.PageSetting{},
		Options: struct {
			WithoutCount  bool
			Scenes        string
			CustomDBOrder string
			CustomFunc    func() interface{}
		}{WithoutCount: true},
		Fields: []field.Expr{
			cladminquery.Q.SysMenu.ALL,
		},
		Conditions: []gen.Condition{},
		Orders: []field.Expr{
			cladminquery.Q.SysMenu.ParentID,
			cladminquery.Q.SysMenu.OrderNum,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, info)
}

// Select
// @Description: 上级菜单type 为0,1类型
// @param c
func Select(c *gin.Context) {
	menuService := menuservice.NewMenuService(c)
	info, _, errNo := menuService.InfoList(&service.ListParams{
		PS: util.PageSetting{},
		Options: struct {
			WithoutCount  bool
			Scenes        string
			CustomDBOrder string
			CustomFunc    func() interface{}
		}{
			WithoutCount: true,
		},
		Fields: []field.Expr{
			cladminquery.Q.SysMenu.ALL,
		},
		Conditions: []gen.Condition{
			cladminquery.Q.SysMenu.Type.Neq(2),
		},
		Orders: []field.Expr{
			cladminquery.Q.SysMenu.ParentID,
			cladminquery.Q.SysMenu.OrderNum,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, info)
}
