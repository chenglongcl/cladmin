package menu

import (
	. "cladmin/handler"
	"cladmin/model"
	"cladmin/pkg/errno"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var (
		r CreateRequest
		m model.Menu
	)
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	util.StructCopy(&m, &r)
	if err := m.CreateMenu(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
