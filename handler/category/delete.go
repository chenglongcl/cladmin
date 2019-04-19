package category

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/category_service"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	categoryService := category_service.Category{
		Id: r.Id,
	}
	if errNo := categoryService.Delete(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
