package oss_service

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"github.com/json-iterator/go"
)

type OssConfig struct {
	AliYunAccessKeyId     string `json:"aliyunAccessKeyId"`
	AliYunAccessKeySecret string `json:"aliyunAccessKeySecret"`
	AliYunBucketName      string `json:"aliyunBucketName"`
	AliYunEndPoint        string `json:"aliyunEndPoint"`
	OssType               string `json:"ossType"`
}

func (a *OssConfig) SaveConfig() *errno.Errno {
	paramValue, _ := jsoniter.MarshalToString(a)
	data := map[string]interface{}{
		"id":          1,
		"param_value": paramValue,
	}
	if err := model.EditConfig(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}
