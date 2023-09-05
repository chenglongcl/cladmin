package article

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/articleservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	articleService := articleservice.NewArticleService(c)
	articleService.UserID = r.UserID
	articleService.CateID = r.CateID
	articleService.Title = r.Title
	articleService.Thumb = r.Thumb
	articleService.Content = r.Content
	if _, errNo := articleService.Add(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
