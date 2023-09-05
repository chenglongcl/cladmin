package user

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"cladmin/service/userservice"
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
	userService := userservice.NewUserService(c)
	info, count, errNo := userService.InfoList(&service.ListParams{
		PS: ps,
		Fields: []field.Expr{
			cladminquery.Q.SysUser.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.UserName != "" {
				conditions = append(conditions, cladminquery.Q.SysUser.Username.Like(util.StringBuilder("%", r.UserName, "%")))
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysUser.ID.Desc(),
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, info))
}
