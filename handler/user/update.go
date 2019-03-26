package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service"
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
	userService := service.User{
		Id:       r.Id,
		Password: r.Password,
		Mobile:   r.Mobile,
		Email:    r.Email,
		Status:   r.Status,
		RoleId:   r.RoleId,
	}
	errNo := userService.Edit()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	inject.Obj.Common.UserAPI.LoadPolicy(userService.Id)
	SendResponse(c, nil, nil)
}
