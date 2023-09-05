package category

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/categoryservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	categoryService := categoryservice.NewCategoryService(c)
	categoryService.ParentID = r.ParentID
	categoryService.Name = r.Name
	categoryService.Icon = r.Icon
	categoryService.OrderNum = r.OrderNum
	if _, errNo := categoryService.Add(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
