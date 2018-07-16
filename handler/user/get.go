package user

import (
	"github.com/gin-gonic/gin"
	"apiserver/model"
	. "apiserver/handler"
	"apiserver/pkg/errno"
)

func Get(c *gin.Context) {
	userId, _ := c.Get("userId")
	fields := []string{
		"id",
		"username",
		"mobile",
	}
	user, err := model.GetUser(userId.(uint64), fields)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}
