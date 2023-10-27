package dept

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"cladmin/service/deptservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func List(c *gin.Context) {
	var (
		r ListRequest
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, nil, errno.ErrBind)
		return
	}
	deptService := deptservice.NewDeptService(c)
	infos, errNo := deptService.Tree(&service.ListParams{
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
			cladminquery.Q.SysDept.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysDept.ParentID,
			cladminquery.Q.SysDept.Sort,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, infos)
}
