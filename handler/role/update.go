package role

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/role_service"
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
	roleService := role_service.Role{
		Id:         r.Id,
		RoleName:   r.RoleName,
		Remark:     r.Remark,
		MenuIdList: r.MenuIdList,
	}
	if errNo := roleService.Edit(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	inject.Obj.Common.RoleAPI.LoadPolicy(roleService.Id)
	SendResponse(c, nil, nil)
}
