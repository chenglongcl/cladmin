package oss_service

import (
	"bytes"
	"cladmin/pkg/errno"
	"cladmin/util"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

type Oss struct {
	AliYunOssClient *oss.Client `inject:""`
}

func (o *Oss) PutObjectWithByte(file multipart.File, header *multipart.FileHeader) (string, *errno.Errno) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", nil
	}
	// 获取存储空间。
	bucket, err := o.AliYunOssClient.Bucket("aisyweixinpic")
	if err != nil {
		return "", errno.ErrAliYunBucket
	}
	newFileName, _ := util.GenStr(16)
	objectKey := time.Now().Format("20060102") + "/" + newFileName +
		strings.ToLower(path.Ext(header.Filename))
	var fileUrl string
	finished := make(chan bool, 1)
	go func() {
		bucketInfo, _ := o.AliYunOssClient.GetBucketInfo("aisyweixinpic")
		fileUrl = "http://" + bucketInfo.BucketInfo.Name + "." +
			bucketInfo.BucketInfo.ExtranetEndpoint + "/" + objectKey
		close(finished)
	}()
	// 上传Byte数组。
	err = bucket.PutObject(objectKey, bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", errno.ErrAliYunOssUploadFail
	}
	<-finished
	return fileUrl, nil
}
