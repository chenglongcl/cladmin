package uploadservice

import (
	"cladmin/pkg/errno"
	"cladmin/service"
	"github.com/duke-git/lancet/v2/random"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type upload struct {
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Upload = *upload

func NewUploadService(ctx *gin.Context, opts ...service.Option) Upload {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &upload{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Upload) UploadFile(file *multipart.FileHeader) (string, string, *errno.Errno) {
	// Get the file suffix name
	fileSuffix := strings.ToLower(path.Ext(file.Filename))
	// Rename filename and set savePath
	savePath := viper.GetString("upload_path")
	// Set Folder by date
	date := time.Now().Format("20060102")
	//Folder isNotExist
	if _, err := os.Stat(savePath + date); err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(savePath+date, os.ModePerm)
	}
	//Set saveFileName
	saveFileName := random.RandNumeralOrLetter(16)
	dst := savePath + date + "/" + saveFileName + fileSuffix
	if err := a.ctx.SaveUploadedFile(file, dst); err != nil {
		return "", "", errno.ErrUploadFail
	}
	return dst, saveFileName + fileSuffix, nil
}
