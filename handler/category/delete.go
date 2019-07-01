package category

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/categoryservice"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	categoryService := categoryservice.Category{
		ID: r.ID,
	}
	if errNo := categoryService.Delete(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
