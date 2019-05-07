package oss

import (
	. "cladmin/handler"
	"cladmin/router/middleware/inject"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebUploadSign(c *gin.Context) {
	sign, errNo := inject.Obj.Common.AliYunOssApi.WebUploadSign()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	c.JSON(http.StatusOK, sign)
}
