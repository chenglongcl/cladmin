package article

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"cladmin/service/articleservice"
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
	articleService := articleservice.NewArticleService(c)
	articleService.CateID = r.CateID
	articleService.Title = r.Title
	info, count, errNo := articleService.InfoList(&service.ListParams{
		PS: ps,
		Fields: []field.Expr{
			cladminquery.Q.SysArticle.ID, cladminquery.Q.SysArticle.UserID, cladminquery.Q.SysArticle.CateID,
			cladminquery.Q.SysArticle.Title, cladminquery.Q.SysArticle.Thumb, cladminquery.Q.SysArticle.Content,
			cladminquery.Q.SysArticle.ReleaseTime,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.CateID != 0 {
				conditions = append(conditions, cladminquery.Q.SysArticle.CateID.Eq(r.CateID))
			}
			if r.Title != "" {
				conditions = append(conditions, cladminquery.Q.SysArticle.Title.Like(util.StringBuilder("%", r.Title, "%")))
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysArticle.ID.Desc(),
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, info))
}
