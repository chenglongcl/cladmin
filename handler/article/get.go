package article

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/article_service"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
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
	articleService := article_service.Article{
		Id: r.Id,
	}
	article, errNo := articleService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	var thumb []string
	jsoniter.UnmarshalFromString(article.Thumb, &thumb)
	SendResponse(c, nil, GetResponse{
		Id:          article.Id,
		CateId:      article.CateId,
		Title:       article.Title,
		Content:     article.Content,
		Thumb:       article.Thumb,
		ReleaseTime: article.ReleaseTime,
	})
}
