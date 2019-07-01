package category

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/categoryservice"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	categoryService := categoryservice.Category{
		ID: r.ID,
	}
	category, errNo := categoryService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		ID:         category.ID,
		ParentID:   category.ParentID,
		Name:       category.Name,
		Icon:       category.Icon,
		OrderNum:   category.OrderNum,
		CreateTime: category.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
