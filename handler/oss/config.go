package oss

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/pkg/oss"
	"cladmin/service/configservice"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"gorm.io/gen"
)

func SaveConfig(c *gin.Context) {
	var r SaveConfigRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	paramValue, _ := jsoniter.MarshalToString(r)
	configService := configservice.NewConfigService(c)
	if errNo := configService.Edit([]gen.Condition{
		cladminquery.Q.SysConfig.ID.Eq(1),
	}, map[string]interface{}{
		"param_value": paramValue,
	}); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	_ = oss.SelectClient("ali").ResetClient()
	handler.SendResponse(c, nil, nil)
}
