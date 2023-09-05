package role

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"cladmin/service/roleservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func List(c *gin.Context) {
	var (
		r  ListRequest
		ps util.PageSetting
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ps.Setting(r.Page, r.Limit)
	roleService := roleservice.NewRoleService(c)
	roleInfos, count, errNo := roleService.InfoList(&service.ListParams{
		PS: ps,
		Fields: []field.Expr{
			cladminquery.Q.SysRole.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.RoleName != "" {
				conditions = append(conditions, cladminquery.Q.SysRole.RoleName.Like(util.StringBuilder("%", r.RoleName, "%")))
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysRole.ID,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, roleInfos))
}

func Select(c *gin.Context) {
	roleService := roleservice.NewRoleService(c)
	roleInfos, _, errNo := roleService.InfoList(&service.ListParams{
		PS: util.PageSetting{},
		Options: struct {
			WithoutCount  bool
			Scenes        string
			CustomDBOrder string
			CustomFunc    func() interface{}
		}{WithoutCount: true},
		Fields: []field.Expr{
			cladminquery.Q.SysRole.ALL,
		},
		Conditions: nil,
		Orders: []field.Expr{
			cladminquery.Q.SysRole.ID,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, roleInfos)
}
