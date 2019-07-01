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
		ID: r.ID,
	}
	user, errNo := userService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	roleIDList := make([]uint64, len(user.Role))
	for _, role := range user.Role {
		roleIDList = append(roleIDList, role.ID)
	}
	SendResponse(c, nil, GetResponse{
		UserID:       user.ID,
		Username:     user.Username,
		CreateTime:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		CreateUserID: user.CreateUserID,
		Email:        user.Email,
		Mobile:       user.Mobile,
		Status:       user.Status,
		RoleIDList:   roleIDList,
	})
}

func GetPersonalInfo(c *gin.Context) {
	id, _ := c.Get("userID")
	userService := userservice.User{
		ID: id.(uint64),
	}
	user, errNo := userService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		UserID:       user.ID,
		Username:     user.Username,
		CreateTime:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		CreateUserID: user.CreateUserID,
		Email:        user.Email,
		Mobile:       user.Mobile,
		Status:       user.Status,
		RoleIDList:   nil,
	})
}
