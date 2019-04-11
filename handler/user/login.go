package user

import (
	. "cladmin/handler"
	"cladmin/model"
	"cladmin/pkg/auth"
	"cladmin/pkg/errno"
	"cladmin/pkg/token"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.User
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	//Compare the login password with user password
	if err := auth.Compare(user.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	//user locked
	if user.Status != 1 {
		SendResponse(c, errno.ErrDisabledUser, nil)
		return
	}
	// Sign the json web token.
	t, e, re, _ := token.Sign(c, token.Context{ID: user.Id, Username: user.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	SendResponse(c, nil, CreateResponse{
		Username:         user.Username,
		Token:            t,
		ExpiredAt:        e,
		RefreshExpiredAt: re,
	})
}
