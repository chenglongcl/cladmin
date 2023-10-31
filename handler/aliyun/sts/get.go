package sts

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/aliyun"
	"cladmin/pkg/aliyun/sts"
	"cladmin/pkg/errno"
	"cladmin/service/configservice"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"net/http"
)

func GetAssumeRole(c *gin.Context) {
	var (
		resp      GetTokenResponse
		stsConfig sts.Config
	)
	configService := configservice.NewConfigService(c)
	configModel, errNo := configService.Get([]field.Expr{
		cladminquery.Q.SysConfig.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysConfig.ParamKey.Eq("ALI_STS_CONFIG_KEY"),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if configModel == nil || configModel.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	_ = jsoniter.UnmarshalFromString(configModel.ParamValue, &stsConfig)
	request := &sts20150401.AssumeRoleRequest{}
	request.SetDurationSeconds(3600)
	request.SetRoleArn(stsConfig.AliYunRoleArn)
	request.SetRoleSessionName(stsConfig.AliYunRoleSessionName)
	result, err := aliyun.GetAliYunOpenApiClients().STS.Client.AssumeRole(request)
	if err != nil || tea.Int32Value(result.StatusCode) != http.StatusOK {
		handler.SendResponse(c, &errno.Errno{
			Code:    29999,
			Message: "获取阿里云STS AssumeRole失败",
		}, nil)
		return
	}
	resp = GetTokenResponse{
		StatusCode:      tea.Int32Value(result.StatusCode),
		AccessKeyId:     tea.StringValue(result.Body.Credentials.AccessKeyId),
		AccessKeySecret: tea.StringValue(result.Body.Credentials.AccessKeySecret),
		SecurityToken:   tea.StringValue(result.Body.Credentials.SecurityToken),
		Expiration:      tea.StringValue(result.Body.Credentials.Expiration),
		Region:          stsConfig.AliYunOSSRegion,
		Bucket:          stsConfig.AliYunOSSBucket,
	}
	handler.SendResponse(c, nil, resp)
}
