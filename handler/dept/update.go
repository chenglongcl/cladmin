package dept

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/deptservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	deptService := deptservice.NewDeptService(c)
	updateData := make(map[string]interface{})
	updateData["parent_id"] = r.ParentID
	updateData["name"] = r.Name
	updateData["sort"] = r.Sort
	if errNo := deptService.Edit([]gen.Condition{
		cladminquery.Q.SysDept.ID.Eq(r.ID),
	}, updateData); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
