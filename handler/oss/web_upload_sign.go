package oss

import (
	. "cladmin/handler"
	"cladmin/pkg/oss"
	"cladmin/service/oss_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebUploadSign(c *gin.Context) {
	aliyunOssService := &oss_service.AliyunOss{
		Client: oss.MyOss.SelectAliyun(),
	}
	sign, errNo := aliyunOssService.WebUploadSign()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	c.JSON(http.StatusOK, sign)
}
