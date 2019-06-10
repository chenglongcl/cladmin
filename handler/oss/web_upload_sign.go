package oss

import (
	. "cladmin/handler"
	"cladmin/pkg/oss/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebUploadSign(c *gin.Context) {
	sign, errNo := client.DefaultAliClient().WebUploadSign()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	c.JSON(http.StatusOK, sign)
}
