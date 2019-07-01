package oss

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/pkg/oss"
	"cladmin/service/ossservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func SaveConfing(c *gin.Context) {
	var r SaveConfigRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	ossService := ossservice.OssConfig{}
	util.StructCopy(&ossService, &r)
	if errNo := ossService.SaveConfig(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	if ossService.AliYunEndPoint != "" &&
		ossService.AliYunAccessKeyID != "" &&
		ossService.AliYunAccessKeySecret != "" {
		oss.SelectClient("ali").ResetClient()
	}
	SendResponse(c, nil, nil)
}
