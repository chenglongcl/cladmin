package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/userservice"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	userService := userservice.User{
		Id: r.Id,
	}
	user, errNo := userService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	roleIdList := make([]uint64, len(user.Role))
	for _, role := range user.Role {
		roleIdList = append(roleIdList, role.Id)
	}
	SendResponse(c, nil, GetResponse{
		UserId:       user.Id,
		Username:     user.Username,
		CreateTime:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		CreateUserId: user.CreateUserId,
		Email:        user.Email,
		Mobile:       user.Mobile,
		Status:       user.Status,
		RoleIdList:   roleIdList,
	})
}

func GetPersonalInfo(c *gin.Context) {
	id, _ := c.Get("userId")
	userService := userservice.User{
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
