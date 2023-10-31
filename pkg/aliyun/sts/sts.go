package sts

import (
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"context"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/chenglongcl/log"
	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	AliYunAccessKeyID     string `json:"aliyunAccessKeyId"`
	AliYunAccessKeySecret string `json:"aliyunAccessKeySecret"`
	AliYunEndPoint        string `json:"aliyunEndPoint"`
	AliYunRoleArn         string `json:"aliyunRoleArn"`
	AliYunRoleSessionName string `json:"aliyunRoleSessionName"`
	AliYunOSSRegion       string `json:"aliyunOSSRegion"`
	AliYunOSSBucket       string `json:"aliyunOSSBucket"`
}

type STS struct {
	Client *sts20150401.Client
}

func (a *STS) NewClient() error {
	var (
		err         error
		configModel *cladminmodel.SysConfig
		stsConfig   Config
		stsClient   *sts20150401.Client
	)
	configModel, err = cladminquery.Q.WithContext(context.Background()).SysConfig.Where(
		cladminquery.Q.SysConfig.ParamKey.Eq("ALI_STS_CONFIG_KEY"),
	).Take()
	if err != nil {
		return err
	}
	if configModel == nil || configModel.ID == 0 {
		return errors.New("阿里云STS配置不存在")
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
		return err
	}
	a.Client = stsClient
	return nil
}

func (a *STS) ResetClient() error {
	if err := a.NewClient(); err != nil {
		log.Errorf(err, "阿里云STS客户端重置失败")
		return err
	}
	return nil
}
