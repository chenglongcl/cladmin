package category

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/categoryservice"
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
	categoryService := categoryservice.Category{
		Id:       r.Id,
		ParentId: r.ParentId,
		Name:     r.Name,
		Icon:     r.Icon,
		OrderNum: r.OrderNum,
	}
	if errNo := categoryService.Edit(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
