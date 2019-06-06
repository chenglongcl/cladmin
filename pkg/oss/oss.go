package oss

import (
	"cladmin/model"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/json-iterator/go"
)

var MyOss *ClAdminOss

type ClAdminOss struct {
	aliyun *oss.Client
}

func Init() {
	MyOss = &ClAdminOss{
		aliyun: initAliyun(),
	}
}

func initAliyun() *oss.Client {
	var aliyunClient *oss.Client
	aliyunConfig := make(map[string]string, 0)
	aliyunConfigStr, _ := model.GetConfigByParamKey("CLOUD_STORAGE_CONFIG_KEY")
	jsoniter.UnmarshalFromString(aliyunConfigStr.ParamValue, &aliyunConfig)
	if aliyunConfig["aliyunEndPoint"] != "" &&
		aliyunConfig["aliyunAccessKeyId"] != "" &&
		aliyunConfig["aliyunAccessKeySecret"] != "" {
		aliyunClient, _ = oss.New(aliyunConfig["aliyunEndPoint"],
			aliyunConfig["aliyunAccessKeyId"],
			aliyunConfig["aliyunAccessKeySecret"])
	}
	return aliyunClient
}

func (o *ClAdminOss) SelectAliyun() *oss.Client {
	return o.aliyun
}

func (o *ClAdminOss) ResetAliyun() bool {
	var aliyunClient *oss.Client
	aliyunConfig := make(map[string]string, 0)
	aliyunConfigStr, _ := model.GetConfigByParamKey("CLOUD_STORAGE_CONFIG_KEY")
	jsoniter.UnmarshalFromString(aliyunConfigStr.ParamValue, &aliyunConfig)
	if aliyunConfig["aliyunEndPoint"] != "" &&
		aliyunConfig["aliyunAccessKeyId"] != "" &&
		aliyunConfig["aliyunAccessKeySecret"] != "" {
		aliyunClient, _ = oss.New(aliyunConfig["aliyunEndPoint"],
			aliyunConfig["aliyunAccessKeyId"],
			aliyunConfig["aliyunAccessKeySecret"])
		o.aliyun = aliyunClient
		return true
	} else {
		return false
	}
}
