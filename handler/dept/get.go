package dept

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/deptservice"
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
	deptService := deptservice.NewDeptService(c)
	deptModel, errNo := deptService.Get([]field.Expr{
		cladminquery.Q.SysDept.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysDept.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if deptModel == nil || deptModel.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, GetResponse{
		ID:       deptModel.ID,
		ParentID: deptModel.ParentID,
		Name:     deptModel.Name,
		Sort:     deptModel.Sort,
	})
}
