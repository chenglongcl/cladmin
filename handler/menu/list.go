package menu

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"cladmin/service/menuservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"github.com/kakuilan/kgo"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"strings"
)

func List(c *gin.Context) {
	var (
		r ListRequest
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, nil, errno.ErrBind)
		return
	}
	menuService := menuservice.NewMenuService(c)
	infos, errNo := menuService.Tree(&service.ListParams{
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
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.MenuTypes != "" {
				strArr := strings.Split(r.MenuTypes, ",")
				if len(strArr) > 0 {
					typeArr := make([]int64, len(strArr))
					for i, s := range strArr {
						typeArr[i] = kgo.KConv.Str2Int64(s)
					}
					conditions = append(conditions, cladminquery.Q.SysMenu.Type.In(typeArr...))
				}
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysMenu.ParentID,
			cladminquery.Q.SysMenu.OrderNum,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, infos)
}
