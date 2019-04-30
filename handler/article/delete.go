package article

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/article_service"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	articleService := article_service.Article{
		Id: r.Id,
	}
	if errNo := articleService.Delete(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
