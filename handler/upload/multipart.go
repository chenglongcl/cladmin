package upload

import (
	"bytes"
	"cladmin/pkg/constvar"
	"cladmin/pkg/redisgo"
	redsync2 "cladmin/pkg/redsync"
	"cladmin/util"
	"fmt"
	"github.com/duke-git/lancet/v2/algorithm"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"github.com/go-redsync/redsync/v4"
	"github.com/golang-module/carbon/v2"
	"github.com/jinzhu/copier"
	"github.com/kakuilan/kgo"
	"github.com/spf13/viper"
	"go.jetpack.io/typeid"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"strings"
)

func InitOrCompleteMultipartUpload(c *gin.Context) {
	if uploadId := c.Query("uploadId"); uploadId != "" {
		completeMultipartUpload(c)
	} else {
		initiateMultipartUpload(c)
	}
}

func initiateMultipartUpload(c *gin.Context) {
	var (
		resp InitiateMultipartUploadResponse
		err  error
	)
	uploadPath := viper.GetString("upload_path")
	fileDomain := viper.GetString("file_domain")
	objectName := c.Param("objectName")
	if uploadPath == "" || !strutil.HasPrefixAny(objectName, []string{"/"}) {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	tid, _ := typeid.New("")
	objectInfo := ObjectInfo{
		Key:        strutil.After(objectName, "/"),
		UploadId:   strings.ToUpper(tid.String()),
		RequestUrl: util.StringBuilder(fileDomain, "/multipartUpload", objectName),
		Location:   util.StringBuilder(fileDomain, "/", uploadPath, strutil.After(objectName, "/")),
	}
	if err = redisgo.My().Set(
		util.StringBuilder(constvar.InitiateMultipartUploadRedisKey, objectInfo.UploadId),
		objectInfo, 86400); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	if err = copier.Copy(&resp, &objectInfo); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	c.XML(http.StatusOK, resp)
}

func MultipartUploadPart(c *gin.Context) {
	var (
		r          MultipartUploadPartRequest
		objectInfo ObjectInfo
		err        error
	)
	objectKey := strutil.After(c.Param("objectName"), "/")
	if err = c.BindQuery(&r); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	if err = redisgo.My().GetObject(
		util.StringBuilder(constvar.InitiateMultipartUploadRedisKey, r.UploadId), &objectInfo); err != nil || objectKey != objectInfo.Key {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	uploadPath := viper.GetString("upload_path")
	if uploadPath == "" {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	chunkBuffer := &bytes.Buffer{}
	if _, err = io.Copy(chunkBuffer, c.Request.Body); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	//创建目录
	chunkDir := util.StringBuilder(uploadPath, "chunk/", r.UploadId, "/")
	if err = fileutil.CreateDir(chunkDir); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	//创建文件
	chunkPath := util.StringBuilder(chunkDir, convertor.ToString(r.PartNumber))
	if ok := fileutil.CreateFile(chunkPath); !ok {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	//写入文件
	err = fileutil.WriteBytesToFile(chunkPath, chunkBuffer.Bytes())
	if err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	md5Str, err := cryptor.Md5File(chunkPath)
	if err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	md5StrUpper := strings.ToUpper(md5Str)
	c.Header("Etag", fmt.Sprintf("\"%s\"", md5StrUpper))
	c.Header("Access-Control-Expose-Headers", "Etag")
	c.XML(http.StatusOK, nil)
}

func completeMultipartUpload(c *gin.Context) {
	var (
		r          CompleteMultipartUploadRequest
		resp       CompleteMultipartUploadResponse
		objectInfo ObjectInfo
		err        error
	)
	if err = c.BindXML(&r.CompleteMultipartUpload); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	objectKey := strutil.After(c.Param("objectName"), "/")
	uploadId := c.Query("uploadId")
	if err = redisgo.My().GetObject(
		util.StringBuilder(constvar.InitiateMultipartUploadRedisKey, uploadId), &objectInfo); err != nil || objectKey != objectInfo.Key {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	rs := redsync.New(redsync2.GetPool())
	mutexName := util.StringBuilder("complete-multipart-mutex-", uploadId)
	mutex := rs.NewMutex(mutexName)
	defer func() {
		_, _ = mutex.Unlock()
	}()
	if err = mutex.Lock(); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	uploadPath := viper.GetString("upload_path")
	if uploadPath == "" {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	//创建合并后文件目录
	fileDir := util.StringBuilder(uploadPath, kgo.KFile.Dirname(objectKey), "/")
	if err = fileutil.CreateDir(fileDir); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	//创建合并后文件
	filePath := util.StringBuilder(fileDir, kgo.KFile.Basename(objectKey))
	if ok := fileutil.CreateFile(filePath); !ok {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	//切片文件目录
	chunkDir := util.StringBuilder(uploadPath, "chunk/", uploadId, "/")
	for _, part := range r.CompleteMultipartUpload.Parts {
		//切片文件路径
		chunkPath := util.StringBuilder(chunkDir, convertor.ToString(part.PartNumber))
		//切片文件是否存在
		if isFileExist := fileutil.IsExist(chunkPath); !isFileExist {
			continue
		}
		//验证md5
		md5Str, err := cryptor.Md5File(chunkPath)
		if err != nil {
			continue
		}
		if strings.ToUpper(md5Str) != part.ETag {
			continue
		}
		//
		if fileBuffer, err := kgo.KFile.ReadFile(chunkPath); err == nil {
			if err := kgo.KFile.AppendFile(filePath, fileBuffer); err != nil {
				continue
			}
		}
	}
	//删除切片文件、删除redis记录
	_ = kgo.KFile.DelDir(chunkDir, true)
	_ = redisgo.My().Del(util.StringBuilder(constvar.InitiateMultipartUploadRedisKey, uploadId))
	//
	resp.CompleteMultipartUploadResult.Location = objectInfo.Location
	resp.CompleteMultipartUploadResult.Key = objectInfo.Key
	resp.CompleteMultipartUploadResult.ETag, _ = cryptor.Md5File(filePath)
	c.XML(http.StatusOK, resp.CompleteMultipartUploadResult)
}

func ListParts(c *gin.Context) {
	var (
		resp       ListPartsResponse
		objectInfo ObjectInfo
		err        error
	)
	objectKey := strutil.After(c.Param("objectName"), "/")
	uploadId := c.Query("uploadId")
	if err = redisgo.My().GetObject(
		util.StringBuilder(constvar.InitiateMultipartUploadRedisKey, uploadId), &objectInfo); err != nil || objectKey != objectInfo.Key {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	resp.ListPartsResult.Key = objectKey
	resp.ListPartsResult.UploadId = uploadId
	uploadPath := viper.GetString("upload_path")
	if uploadPath == "" {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	chunkDir := util.StringBuilder(uploadPath, "chunk/", uploadId, "/")
	errGroup := &errgroup.Group{}
	chunkNames, _ := fileutil.ListFileNames(chunkDir)
	for _, cn := range chunkNames {
		chunkName := cn
		errGroup.Go(func() error {
			chunkPath := util.StringBuilder(chunkDir, chunkName)
			md5Str, err := cryptor.Md5File(chunkPath)
			if err != nil {
				return err
			}
			mtTime, err := fileutil.MTime(chunkPath)
			if err != nil {
				return err
			}
			size, err := fileutil.FileSize(chunkPath)
			if err != nil {
				return err
			}
			resp.ListPartsResult.Part = append(resp.ListPartsResult.Part, Part{
				PartNumber:   kgo.KConv.Str2Uint64(chunkName),
				LastModified: carbon.CreateFromTimestamp(mtTime).ToDateTimeString(),
				ETag:         md5Str,
				Size:         size,
			})
			return nil
		})

	}
	if err = errGroup.Wait(); err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}
	algorithm.InsertionSort(resp.ListPartsResult.Part, &partsComparator{})
	c.XML(http.StatusOK, resp.ListPartsResult)
}

type partsComparator struct{}

func (pc *partsComparator) Compare(v1 any, v2 any) int {
	p1, _ := v1.(Part)
	p2, _ := v2.(Part)

	//ascending order
	if p1.PartNumber < p2.PartNumber {
		return -1
	} else if p1.PartNumber > p2.PartNumber {
		return 1
	}
	return 0
}
