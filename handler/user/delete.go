package user

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/userservice"
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
			userService := userservice.NewUserService(c)
			userModel, errNo := userService.Get([]field.Expr{
				cladminquery.Q.SysUser.ALL,
			}, []gen.Condition{
				cladminquery.Q.SysUser.ID.Eq(id),
			})
			if errNo != nil {
				return errNo
			}
			if userModel == nil || userModel.ID == 0 {
				return errno.ErrRecordNotFound
			}
			if errNo = userService.Delete(userModel); errNo != nil {
				return errNo
			}
			inject.Obj.Enforcer.DeleteUser(userModel.Username)
			return nil
		})
	}
	if errNo := errGroup.Wait(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
