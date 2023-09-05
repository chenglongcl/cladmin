package article

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/articleservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	articleService := articleservice.NewArticleService(c)
	articleModel, errNo := articleService.Get([]field.Expr{
		cladminquery.Q.SysArticle.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysArticle.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if articleModel == nil || articleModel.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, GetResponse{
		ID:          articleModel.ID,
		CateID:      articleModel.CateID,
		Title:       articleModel.Title,
		Content:     articleModel.Content,
		Thumb:       articleModel.Thumb,
		ReleaseTime: articleModel.ReleaseTime,
	})
}
