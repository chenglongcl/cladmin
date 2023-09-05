package role

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/roleservice"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	errGroup := &errgroup.Group{}
	for _, _id := range r.Ids {
		id := _id
		errGroup.Go(func() error {
			roleService := roleservice.NewRoleService(c)
			roleModel, errNo := roleService.Get([]field.Expr{
				cladminquery.Q.SysRole.ALL,
			}, []gen.Condition{
				cladminquery.Q.SysRole.ID.Eq(id),
			})
			if errNo != nil {
				return errNo
			}
			if roleModel == nil || roleModel.ID == 0 {
				return errno.ErrRecordNotFound
			}
			if errNo = roleService.Delete(roleModel); errNo != nil {
				return errNo
			}
			inject.Obj.Enforcer.DeleteRole(roleModel.RoleName)
			return nil
		})
	}
	if errNo := errGroup.Wait(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
