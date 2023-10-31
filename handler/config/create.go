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

func Create(c *gin.Context) {
	var (
		r CreateRequest
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
	//2.新增配置
	configService.ParamKey = r.ParamKey
	configService.ParamValue = r.ParamValue
	configService.Type = r.Type
	configService.Status = r.Status
	configService.Remark = r.Remark
	if configModel, errNo = configService.Add(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	go func() {
		switch configModel.ParamKey {
		case "CLOUD_STORAGE_ALI_CONFIG_KEY":
			_ = cloudstorage.GetCloudStorage().AliYun.ResetClient()
		case "ALI_STS_CONFIG_KEY":
			_ = aliyun.GetAliYunOpenApiClients().STS.ResetClient()
		}
	}()
	handler.SendResponse(c, nil, nil)
}
