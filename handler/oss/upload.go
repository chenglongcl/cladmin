package oss

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/pkg/oss"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	var r UploadOssRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		handler.SendResponse(c, errno.ErrUploadFile, nil)
		return
	}
	switch r.OSSName {
	case "aliYunOSS":
		fileUrl, errNo := oss.SelectClient("ali").UpLoad(file, header)
		if errNo != nil {
			handler.SendResponse(c, errNo, nil)
			return
		}
		handler.SendResponse(c, nil, fileUrl)
	}
}
