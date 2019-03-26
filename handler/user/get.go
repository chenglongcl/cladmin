package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := service.User{
		Id: r.Id,
	}
	user, errNo := userService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, user)
}
