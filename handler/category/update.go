package category

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/categoryservice"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	categoryService := categoryservice.Category{
		ID:       r.ID,
		ParentID: r.ParentID,
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
