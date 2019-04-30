package article

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/article_service"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	articleService := article_service.Article{
		Id:      r.Id,
		UserId:  r.UserId,
		CateId:  r.CateId,
		Title:   r.Title,
		Thumb:   r.Thumb,
		Content: r.Content,
	}
	if errNo := articleService.Edit(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
