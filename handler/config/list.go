package config

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"cladmin/service/configservice"
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
	configService := configservice.NewConfigService(c)
	configInfos, count, errNo := configService.InfoList(&service.ListParams{
		PS: ps,
		Fields: []field.Expr{
			cladminquery.Q.SysConfig.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.ParamKey != "" {
				conditions = append(conditions, cladminquery.Q.SysConfig.ParamKey.Like(util.StringBuilder("%", r.ParamKey, "%")))
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysConfig.ID,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, configInfos))
}
