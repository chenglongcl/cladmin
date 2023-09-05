package bulletin

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"cladmin/service/bulletinservice"
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
	bulletinService := bulletinservice.NewBulletinService(c)
	infos, count, errNo := bulletinService.InfoList(&service.ListParams{
		PS: ps,
		Fields: []field.Expr{
			cladminquery.Q.SysBulletin.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.Title != "" {
				conditions = append(conditions, cladminquery.Q.SysBulletin.Title.Like(util.StringBuilder("%", r.Title, "%")))
			}
			if r.Tag != "" {
				conditions = append(conditions, cladminquery.Q.SysBulletin.Tag.Eq(r.Tag))
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysBulletin.ID.Desc(),
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, infos))
}
