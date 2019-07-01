package article

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/articleservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	articleService := articleservice.Article{
		UserID:  r.UserID,
		CateID:  r.CateID,
		Title:   r.Title,
		Thumb:   r.Thumb,
		Content: r.Content,
	}
	if errNo := articleService.Add(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
