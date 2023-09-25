package config

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/configservice"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	configService := configservice.NewConfigService(c)
	config, errNo := configService.Get([]field.Expr{
		cladminquery.Q.SysConfig.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysConfig.ParamKey.Eq(r.Key),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if config == nil || config.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	gcr := GetCommonResponse{
		ID:       config.ID,
		ParamKey: config.ParamKey,
		Remark:   config.Remark,
		Type:     config.Type,
	}
	switch config.Type {
	case 1:
		handler.SendResponse(c, nil, GetResponseWithOneParam{
			GetCommonResponse: gcr,
			ParamValue:        config.ParamValue,
		})
	case 2:
		paramValue := make(map[string]interface{}, 0)
		_ = jsoniter.UnmarshalFromString(config.ParamValue, &paramValue)
		handler.SendResponse(c, nil, GetResponseWithMultipleParams{
			GetCommonResponse: gcr,
			ParamValue:        paramValue,
		})
	}
}
