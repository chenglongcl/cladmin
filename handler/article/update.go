package article

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/articleservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	articleService := articleservice.NewArticleService(c)
	updateData := make(map[string]interface{})
	updateData["title"] = r.Title
	updateData["thumb"] = r.Thumb
	updateData["content"] = r.Content
	updateData["release_time"] = r.ReleaseTime
	updateData["user_id"] = r.UserID
	updateData["cate_id"] = r.CateID
	if errNo := articleService.Edit([]gen.Condition{
		cladminquery.Q.SysArticle.ID.Eq(r.ID),
	}, updateData); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
