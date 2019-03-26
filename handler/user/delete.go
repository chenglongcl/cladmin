package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := service.User{
		Id: r.Id,
	}
	user, _ := userService.Get()
	if errNo := userService.Delete(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	inject.Obj.Enforcer.DeleteUser(user.Username)
	SendResponse(c, nil, nil)
}
