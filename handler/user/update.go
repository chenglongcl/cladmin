package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/userservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

// Update update a exist user account info.
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
	userService := userservice.User{
		Id:         r.Id,
		Username:   r.Username,
		Password:   r.Password,
		Mobile:     r.Mobile,
		Email:      r.Email,
		Status:     r.Status,
		RoleIdList: r.RoleIdList,
	}
	errNo := userService.Edit()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	inject.Obj.Common.UserAPI.LoadPolicy(userService.Id)
	SendResponse(c, nil, nil)
}

func UpdatePersonal(c *gin.Context) {
	var r UpdatePersonalRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	id, exist := c.Get("userId")
	if !exist {
		SendResponse(c, errno.ErrNotUserExist, nil)
		return
	}
	userService := userservice.User{
		Id:       id.(uint64),
		Password: r.Password,
	}
	errNo := userService.EditPersonal()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
