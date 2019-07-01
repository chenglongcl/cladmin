package config

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/configservice"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	configService := configservice.Config{
		ID:         r.ID,
		ParamKey:   r.ParamKey,
		ParamValue: r.ParamValue,
		Status:     r.Status,
		Type:       r.Type,
		Remark:     r.Remark,
	}
	if errNo := configService.Edit(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
