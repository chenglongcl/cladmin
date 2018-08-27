package user

import (
	"github.com/gin-gonic/gin"
	"apiserver/model"
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/auth"
	"apiserver/pkg/token"
)

func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := model.GetUserByUsername(u.Username, []string{"id", "username", "password"})
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	//Compare the login password with user password
	if err := auth.Compare(user.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
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