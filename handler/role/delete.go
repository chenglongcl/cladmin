package role

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/role_service"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	roleService := role_service.Role{
		Id: r.Id,
	}
	role, _ := roleService.Get()
	if errNo := roleService.Delete(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	inject.Obj.Enforcer.DeleteRole(role.RoleName)
	SendResponse(c, nil, nil)
}
