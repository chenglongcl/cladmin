package user

import (
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/token"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		Mobile:   r.Mobile,
	}
	//Validate the data
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	//Encrypt password
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	if err := u.CreateUser(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	t, e, _ := token.Sign(c, token.Context{ID: u.Id, Username: u.Username}, "")

	rep := CreateResponse{
		Username:  r.Username,
		Token:     t,
		ExpiredAt: e,
	}
	SendResponse(c, nil, rep)
}
