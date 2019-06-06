package oss

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/pkg/oss"
	"cladmin/service/oss_service"
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
		aliyunOssService := &oss_service.AliyunOss{
			Client: oss.MyOss.SelectAliyun(),
		}
		fileUrl, errNo := aliyunOssService.PutObjectWithByte(file, header)
		if errNo != nil {
			SendResponse(c, errNo, nil)
			return
		}
		SendResponse(c, nil, fileUrl)
	}
}
