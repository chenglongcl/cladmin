package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Update update a exist user account info.
func Update(c *gin.Context) {

	userId, _ := c.Get("userId")

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	u.Id = userId.(uint64)

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.UpdateUser(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
