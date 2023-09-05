package category

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/categoryservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	categoryService := categoryservice.NewCategoryService(c)
	updateData := make(map[string]interface{})
	updateData["parent_id"] = r.ParentID
	updateData["name"] = r.Name
	updateData["icon"] = r.Icon
	updateData["order_num"] = r.OrderNum
	if errNo := categoryService.Edit([]gen.Condition{
		cladminquery.Q.SysCategory.ID.Eq(r.ID),
	}, updateData); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
