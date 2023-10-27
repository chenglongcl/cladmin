package aliyun

import (
	"cladmin/dal/cladmindb/cladminquery"
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/chenglongcl/log"
	jsoniter "github.com/json-iterator/go"
)

type STSConfig struct {
	AliYunAccessKeyID     string `json:"aliyunAccessKeyId"`
	AliYunAccessKeySecret string `json:"aliyunAccessKeySecret"`
	AliYunEndPoint        string `json:"aliyunEndPoint"`
	AliYunRoleArn         string `json:"aliyunRoleArn"`
	AliYunRoleSessionName string `json:"aliyunRoleSessionName"`
	AliYunOSSRegion       string `json:"aliyunOSSRegion"`
	AliYunOSSBucket       string `json:"aliyunOSSBucket"`
}

var (
	stsClient *sts20150401.Client
)

func InitAilYunOpenApiClients() {
	//阿里STS
	initSTSClient()
}

func initSTSClient() {
	var (
		err       error
		stsConfig STSConfig
	)
	configModel, err := cladminquery.Q.WithContext(context.Background()).SysConfig.Where(
		cladminquery.Q.SysConfig.ParamKey.Eq("ALI_STS_CONFIG_KEY"),
	).Take()
	if err != nil || configModel == nil || configModel.ID == 0 {
		log.Errorf(err, "获取阿里云STS配置失败")
		return
	}
	_ = jsoniter.UnmarshalFromString(configModel.ParamValue, &stsConfig)
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: tea.String(stsConfig.AliYunAccessKeyID),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(stsConfig.AliYunAccessKeySecret),
	}
	config.Endpoint = tea.String(stsConfig.AliYunEndPoint)
	if stsClient, err = sts20150401.NewClient(config); err != nil {
		log.Errorf(err, "初始阿里STS客户端失败")
	}
}

func ResetClient(clientName string) {
	switch clientName {
	case "sts":
		initSTSClient()
	}
}

func GetAliSTSClient() *sts20150401.Client {
	return stsClient
}
