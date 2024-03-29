package oss

import (
	"cladmin/handler"
	"cladmin/pkg/cloudstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebUploadSign(c *gin.Context) {
	sign, errNo := cloudstorage.GetCloudStorage().AliYun.WebUploadSign()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	c.JSON(http.StatusOK, sign)
}
