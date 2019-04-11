package oss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type AliyunOss struct {
	Client *oss.Client
}

var ALIYUNOSS *AliyunOss

func (a *AliyunOss) Init() {
	ALIYUNOSS = &AliyunOss{}
	ALIYUNOSS.Client, _ = oss.New("http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com",
		"LTAIcpkAxHSr8L6t",
		"tI2UWkXUzutGMtVeIWSzCrX5IKmONS")
}
