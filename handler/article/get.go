package article

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/articleservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	articleService := articleservice.Article{
		ID: r.ID,
	}
	article, errNo := articleService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		ID:          article.ID,
		CateID:      article.CateID,
		Title:       article.Title,
		Content:     article.Content,
		Thumb:       article.Thumb,
		ReleaseTime: article.ReleaseTime,
	})
}
