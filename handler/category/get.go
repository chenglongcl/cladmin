package category

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/categoryservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	categoryService := categoryservice.NewCategoryService(c)
	categoryModel, errNo := categoryService.Get([]field.Expr{
		cladminquery.Q.SysCategory.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysCategory.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, GetResponse{
		ID:         categoryModel.ID,
		ParentID:   categoryModel.ParentID,
		Name:       categoryModel.Name,
		Icon:       categoryModel.Icon,
		OrderNum:   categoryModel.OrderNum,
		CreateTime: categoryModel.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
