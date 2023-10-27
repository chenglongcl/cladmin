package config

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/aliyun"
	"cladmin/pkg/cloudstorage"
	"cladmin/pkg/errno"
	"cladmin/service/configservice"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Update(c *gin.Context) {
	var (
		r UpdateRequest
	)
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if r.Type == 2 && !validator.IsJSON(r.ParamValue) {
		handler.SendResponse(c, &errno.Errno{
			Code:    29999,
			Message: "参数值必须为json格式",
		}, nil)
		return
	}
	configService := configservice.NewConfigService(c)
	//1.查询配置键名是否存在
	configModel, errNo := configService.Get([]field.Expr{
		cladminquery.Q.SysConfig.ID,
	}, []gen.Condition{
		cladminquery.Q.SysConfig.ID.Neq(r.ID),
		cladminquery.Q.SysConfig.ParamKey.Eq(r.ParamKey),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if configModel != nil && configModel.ID != 0 {
		handler.SendResponse(c, &errno.Errno{
			Code:    29999,
			Message: "配置已存在",
		}, nil)
		return
	}
	//2.更新配置
	updateData := make(map[string]interface{})
	updateData["param_key"] = r.ParamKey
	updateData["param_value"] = r.ParamValue
	updateData["status"] = r.Status
	updateData["type"] = r.Type
	updateData["remark"] = r.Remark
	if errNo = configService.Edit([]gen.Condition{
		cladminquery.Q.SysConfig.ID.Eq(r.ID),
		cladminquery.Q.SysConfig.Locked.Is(false),
	}, updateData); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	go func() {
		switch r.ParamKey {
		case "CLOUD_STORAGE_ALI_CONFIG_KEY":
			_ = cloudstorage.SelectClient("ali").ResetClient()
		case "ALI_STS_CONFIG_KEY":
			aliyun.ResetClient("sts")
		}
	}()
	handler.SendResponse(c, nil, nil)
}
