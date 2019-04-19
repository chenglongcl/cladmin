package category

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/category_service"
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
	categoryService := &category_service.Category{
		ParentId: r.ParentId,
		Name:     r.Name,
		Icon:     r.Icon,
		OrderNum: r.OrderNum,
	}
	if errNo := categoryService.Add(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
