package upload

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"github.com/gin-gonic/gin"
)

func ToOss(c *gin.Context) {
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
		fileUrl, errNo := inject.Obj.Common.AliYunOssApi.PutObjectWithByte(file, header);
		if errNo != nil {
			SendResponse(c, errNo, nil)
			return
		}
		SendResponse(c, nil, fileUrl)
	}
}
