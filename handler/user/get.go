package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/user_service"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := user_service.User{
		Id: r.Id,
	}
	user, errNo := userService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, user)
}

func GetPersonalInfo(c *gin.Context) {
	id, _ := c.Get("userId")
	userService := user_service.User{
		Id: id.(uint64),
	}
	user, errNo := userService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		UserId:       user.Id,
		Username:     user.Username,
		CreateTime:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		CreateUserId: user.CreateUserId,
		Email:        user.Email,
		Mobile:       user.Mobile,
		Status:       user.Status,
		RoleIdList:   nil,
	})
}
