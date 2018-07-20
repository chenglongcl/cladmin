package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
	"os"
	"path"
	"apiserver/util"
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"strings"
	"mime"
	"regexp"
)

func Img(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		SendResponse(c, errno.ErrUploadFile, nil)
		return
	}
	// Get the file suffix name
	fileSuffix := strings.ToLower(path.Ext(file.Filename))
	//Validate file
	mimeType := mime.TypeByExtension(fileSuffix)
	if matched, err := regexp.MatchString("^image/[a-zA-Z]+$", mimeType); !matched || err != nil {
		SendResponse(c, errno.ErrUploadMime, nil)
		return
	}
	// Rename filename and set savePath
	savePath := viper.GetString("upload_path.images")
	// Set Folder by date
	date := time.Now().Format("20060102")
	//Folder isNotExist
	if _, err := os.Stat(savePath + date); err != nil && os.IsNotExist(err) {
		os.MkdirAll(savePath+date, os.ModePerm)
	}
	//Set saveFileName
	saveFileName, _ := util.GenShortId()
	dst := savePath + date + "/" + saveFileName + fileSuffix
	if err := c.SaveUploadedFile(file, dst); err != nil {
		SendResponse(c, errno.ErrUploadFail, nil)
		return
	}
	rep := &ImageResponse{
		Path:     dst,
		FileName: saveFileName + fileSuffix,
	}
	SendResponse(c, nil, rep)
}
