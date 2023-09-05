package config

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/configservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	configService := configservice.NewConfigService(c)
	updateData := make(map[string]interface{})
	updateData["param_key"] = r.ParamKey
	updateData["param_value"] = r.ParamValue
	updateData["status"] = r.Status
	updateData["type"] = r.Type
	updateData["remark"] = r.Remark
	if errNo := configService.Edit([]gen.Condition{
		cladminquery.Q.SysConfig.ID.Eq(r.ID),
	}, updateData); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
