package category

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/service"
	"cladmin/service/categoryservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func List(c *gin.Context) {
	categoryService := categoryservice.NewCategoryService(c)
	infos, _, errNo := categoryService.InfoList(&service.ListParams{
		PS: util.PageSetting{},
		Options: struct {
			WithoutCount  bool
			Scenes        string
			CustomDBOrder string
			CustomFunc    func() interface{}
		}{WithoutCount: true},
		Fields: []field.Expr{
			cladminquery.Q.SysCategory.ALL,
		},
		Conditions: []gen.Condition{},
		Orders: []field.Expr{
			cladminquery.Q.SysCategory.ParentID,
			cladminquery.Q.SysCategory.OrderNum,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, infos)
}
