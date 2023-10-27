package dept

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/deptservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	deptService := deptservice.NewDeptService(c)
	//
	childCount, errNo := deptService.Count([]gen.Condition{
		cladminquery.Q.SysDept.ParentID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if childCount > 0 {
		handler.SendResponse(c, &errno.Errno{
			Code:    29999,
			Message: "该部门下存在子部门，无法删除",
		}, nil)
		return
	}
	//
	if errNo = deptService.Delete([]gen.Condition{
		cladminquery.Q.SysDept.ID.Eq(r.ID),
	}); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
