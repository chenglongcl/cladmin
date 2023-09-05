package upload

import (
	"cladmin/handler"
	"cladmin/pkg/errno"
	"github.com/duke-git/lancet/v2/random"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mime"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

func Img(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		handler.SendResponse(c, errno.ErrUploadFile, nil)
		return
	}
	// Get the file suffix name
	fileSuffix := strings.ToLower(path.Ext(file.Filename))
	//Validate file
	mimeType := mime.TypeByExtension(fileSuffix)
	if matched, err := regexp.MatchString("^image/[a-zA-Z]+$", mimeType); !matched || err != nil {
		handler.SendResponse(c, errno.ErrUploadMime, nil)
		return
	}
	// Rename filename and set savePath
	savePath := viper.GetString("upload_path.images")
	// Set Folder by date
	date := time.Now().Format("20060102")
	//Folder isNotExist
	if _, err := os.Stat(savePath + date); err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(savePath+date, os.ModePerm)
	}
	//Set saveFileName
	saveFileName := random.RandNumeralOrLetter(16)
	dst := savePath + date + "/" + saveFileName + fileSuffix
	if err := c.SaveUploadedFile(file, dst); err != nil {
		handler.SendResponse(c, errno.ErrUploadFail, nil)
		return
	}
	rep := &UploadResponse{
		Url:      dst,
		FileName: saveFileName + fileSuffix,
	}
	handler.SendResponse(c, nil, rep)
}
