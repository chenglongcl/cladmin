package role

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/roleservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	roleService := roleservice.NewRoleService(c)
	roleService.RoleName = r.RoleName
	roleService.Remark = r.Remark
	roleService.CreateUserID = r.CreateUserID
	roleService.MenuIDList = r.MenuIDList
	roleModel, errNo := roleService.Add()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	_ = inject.Obj.Common.RoleAPI.LoadPolicy(roleModel.ID)
	handler.SendResponse(c, nil, nil)
}
