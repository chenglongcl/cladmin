package upload

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/uploadservice"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		handler.SendResponse(c, errno.ErrUploadFile, nil)
		return
	}
	uploadService := uploadservice.NewUploadService(c)
	path, fileName, errNo := uploadService.UploadFile(file)
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, UploadResponse{
		Url:      viper.GetString("file_domain") + "/" + path,
		Path:     path,
		FileName: fileName,
	})
}
