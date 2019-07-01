package config

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/configservice"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	configService := configservice.Config{
		ParamKey: r.Key,
	}
	config, errNo := configService.GetByParamKey()
	if errNo != nil {
		SendResponse(c, errNo, nil)
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
		SendResponse(c, nil, GetResponseWithOneParam{
			GetCommonResponse: gcr,
			ParamValue:        config.ParamValue,
		})
	case 2:
		paramValue := make(map[string]interface{}, 0)
		jsoniter.UnmarshalFromString(config.ParamValue, &paramValue)
		SendResponse(c, nil, GetResponseWithMultipleParams{
			GetCommonResponse: gcr,
			ParamValue:        paramValue,
		})
	}
}
