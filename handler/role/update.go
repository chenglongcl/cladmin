package role

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/roleservice"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	roleService := roleservice.NewRoleService(c)
	roleModel, errNo := roleService.Get([]field.Expr{
		cladminquery.Q.SysRole.ID,
	}, []gen.Condition{
		cladminquery.Q.SysRole.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if roleModel == nil || roleModel.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	roleModel.RoleName = r.RoleName
	roleModel.Remark = r.Remark
	roleModel.MenuIDList, _ = jsoniter.MarshalToString(r.MenuIDList)
	if errNo = roleService.Edit(roleModel, r.MenuIDList); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	_ = inject.Obj.Common.RoleAPI.LoadPolicy(roleModel.ID)
	handler.SendResponse(c, nil, nil)
}
