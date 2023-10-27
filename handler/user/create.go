package user

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/pkg/token"
	"cladmin/router/middleware/inject"
	"cladmin/service/userservice"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.NewUserService(c)
	userService.Username = r.Username
	userService.Password = r.Password
	userService.DeptID = r.DeptID
	userService.Mobile = r.Mobile
	userService.Email = r.Email
	userService.Gender = r.Gender
	userService.Status = r.Status
	userService.CreateUserID = r.CreateUserID
	userService.RoleIDList = r.RoleIdList
	userModel, errNo := userService.Add()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	t, e, re, _ := token.Sign(c, token.Context{ID: userModel.ID, Username: userModel.Username, SuperAdmin: userModel.SuperAdmin}, "")
	_ = inject.Obj.Common.UserAPI.LoadPolicy(userModel.ID)
	resp := CreateResponse{
		Username:         r.Username,
		Token:            t,
		ExpiredAt:        e,
		RefreshExpiredAt: re,
	}
	handler.SendResponse(c, nil, resp)
}
