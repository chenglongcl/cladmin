package oss

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/oss_service"
	"cladmin/util"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

func SaveConfing(c *gin.Context) {
	var r SaveConfigRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	ossService := oss_service.OssConfig{}
	util.StructCopy(&ossService, &r)
	if errNo := ossService.SaveConfig(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	inject.Obj.AliYunOssClient, _ = oss.New(ossService.AliYunEndPoint,
		ossService.AliYunAccessKeyId,
		ossService.AliYunAccessKeySecret)
	SendResponse(c, nil, nil)
}
