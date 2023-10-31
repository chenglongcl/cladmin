package aliyun

import (
	"cladmin/pkg/aliyun/sts"
	"github.com/chenglongcl/log"
)

var openApiClients *OpenApiClients

type OpenApiClients struct {
	STS *sts.STS
}

func InitAliYunOpenApiClients() {
	var (
		err error
	)
	openApiClients = &OpenApiClients{}
	//
	stsE := &sts.STS{}
	if err = stsE.NewClient(); err != nil {
		log.Errorf(err, "阿里云STS客户端初始化失败")
	}
	//
	openApiClients = &OpenApiClients{
		STS: stsE,
	}
}

func GetAliYunOpenApiClients() *OpenApiClients {
	return openApiClients
}
