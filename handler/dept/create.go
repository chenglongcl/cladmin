package dept

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/deptservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	deptService := deptservice.NewDeptService(c)
	deptService.ParentID = r.ParentID
	deptService.Name = r.Name
	deptService.Sort = r.Sort
	if _, errNo := deptService.Add(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
