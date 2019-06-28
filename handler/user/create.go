package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/pkg/token"
	"cladmin/router/middleware/inject"
	"cladmin/service/userservice"
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
	userService := userservice.User{
		Username:     r.Username,
		Password:     r.Password,
		Mobile:       r.Mobile,
		Email:        r.Email,
		Status:       r.Status,
		CreateUserId: r.CreateUserId,
		RoleIdList:   r.RoleIdList,
	}
	id, errNo := userService.Add()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	t, e, re, _ := token.Sign(c, token.Context{ID: id, Username: userService.Username}, "")
	inject.Obj.Common.UserAPI.LoadPolicy(id)
	rep := CreateResponse{
		Username:         r.Username,
		Token:            t,
		ExpiredAt:        e,
		RefreshExpiredAt: re,
	}
	SendResponse(c, nil, rep)
}
