package oss

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/pkg/oss"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	var r UploadOssRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		SendResponse(c, errno.ErrUploadFile, nil)
		return
	}
	switch r.OssName {
	case "aliYunOss":
		fileUrl, errNo := oss.SelectClient("ali").UpLoad(file, header)
		if errNo != nil {
			SendResponse(c, errNo, nil)
			return
		}
		SendResponse(c, nil, fileUrl)
	}
}
