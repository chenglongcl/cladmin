package config

import (
	"cladmin/handler"
	"cladmin/pkg/aliyun"
	"cladmin/pkg/errno"
	"cladmin/pkg/oss"
	"cladmin/service/configservice"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func Create(c *gin.Context) {
	var r UpsertRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	configService := configservice.NewConfigService(c)
	if err := copier.Copy(configService, &r); err != nil {
		handler.SendResponse(c, errno.ErrParams, nil)
		return
	}
	updateData := make(map[string]interface{})
	updateData["param_key"] = r.ParamKey
	updateData["param_value"] = r.ParamValue
	updateData["status"] = r.Status
	updateData["type"] = r.Type
	updateData["remark"] = r.Remark
	configModel, errNo := configService.UpsertByID(updateData)
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	go func() {
		switch configModel.ParamKey {
		case "CLOUD_STORAGE_CONFIG_KEY":
			_ = oss.SelectClient("ali").ResetClient()
		case "ALIYUN_STS_CONFIG_KEY":
			aliyun.ResetClient("sts")
		}
	}()
	handler.SendResponse(c, nil, nil)
}
