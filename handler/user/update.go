package user

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/auth"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/userservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

// Update an exist user account info.
func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userModel, errNo := userService.Get([]field.Expr{
		cladminquery.Q.SysUser.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysUser.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if userModel == nil || userModel.ID == 0 {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	userModel.Username = r.Username
	if r.Password != "" {
		userModel.Password, _ = auth.Encrypt(r.Password)
	}
	userModel.DeptID = r.DeptID
	userModel.Mobile = r.Mobile
	userModel.Gender = r.Gender
	userModel.Email = r.Email
	userModel.Status = r.Status
	if errNo = userService.Edit(userModel, r.RoleIDList); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	_ = inject.Obj.Common.UserAPI.LoadPolicy(userModel.ID)
	handler.SendResponse(c, nil, nil)
}

func UpdatePersonal(c *gin.Context) {
	var r UpdatePersonalRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	id := c.GetUint64("userID")
	if id == 0 {
		handler.SendResponse(c, errno.ErrNotUserExist, nil)
		return
	}
	userService := userservice.NewUserService(c)
	updateData := make(map[string]interface{})
	if r.Password != "" {
		updateData["password"], _ = auth.Encrypt(r.Password)
	}
	errNo := userService.EditPersonal([]gen.Condition{
		cladminquery.Q.SysUser.ID.Eq(id),
	}, updateData)
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
