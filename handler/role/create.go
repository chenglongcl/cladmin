package role

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/roleservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	roleService := roleservice.Role{
		RoleName:     r.RoleName,
		Remark:       r.Remark,
		CreateUserId: r.CreateUserId,
		MenuIdList:   r.MenuIdList,
	}
	id, errNo := roleService.Add()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	inject.Obj.Common.RoleAPI.LoadPolicy(id)
	SendResponse(c, nil, nil)
}
