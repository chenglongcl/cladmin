package client

import (
	"bytes"
	"cladmin/model"
	"cladmin/pkg/errno"
	"cladmin/util"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/json-iterator/go"
	"hash"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

type PolicyConfig struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
}

var ali *Ali

type Ali struct {
	Client *oss.Client
}

func InitAliClient() {
	var client *oss.Client
	aliConfig := make(map[string]string, 0)
	aliConfigStr, _ := model.GetConfigByParamKey("CLOUD_STORAGE_CONFIG_KEY")
	jsoniter.UnmarshalFromString(aliConfigStr.ParamValue, &aliConfig)
	if aliConfig["aliyunEndPoint"] != "" &&
		aliConfig["aliyunAccessKeyId"] != "" &&
		aliConfig["aliyunAccessKeySecret"] != "" {
		client, _ = oss.New(aliConfig["aliyunEndPoint"],
			aliConfig["aliyunAccessKeyId"],
			aliConfig["aliyunAccessKeySecret"])
	}
	ali = &Ali{
		Client: client,
	}
}

func DefaultAliClient() *Ali {
	return ali
}

func (o *Ali) UpLoad(file multipart.File, header *multipart.FileHeader) (string, *errno.Errno) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", nil
	}
	if o.Client == nil {
		return "", errno.ErrAliYunOssUploadFail
	}
	ossConfig := make(map[string]string, 0)
	ossConfigStr, _ := model.GetConfigByParamKey("CLOUD_STORAGE_CONFIG_KEY")
	jsoniter.UnmarshalFromString(ossConfigStr.ParamValue, &ossConfig)
	bucketName := ossConfig["aliyunBucketName"]
	// 获取存储空间。
	bucket, err := o.Client.Bucket(bucketName)
	if err != nil || bucket.BucketName == "" {
		return "", errno.ErrAliYunBucket
	}
	newFileName, _ := util.GenStr(16)
	objectKey := time.Now().Format("20060102") + "/" + newFileName +
		strings.ToLower(path.Ext(header.Filename))
	var fileUrl string
	finished := make(chan bool, 1)
	go func() {
		bucketInfo, _ := o.Client.GetBucketInfo(bucketName)
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

func (o *Ali) ResetClient() bool {
	var client *oss.Client
	aliConfig := make(map[string]string, 0)
	aliConfigStr, _ := model.GetConfigByParamKey("CLOUD_STORAGE_CONFIG_KEY")
	jsoniter.UnmarshalFromString(aliConfigStr.ParamValue, &aliConfig)
	if aliConfig["aliyunEndPoint"] != "" &&
		aliConfig["aliyunAccessKeyId"] != "" &&
		aliConfig["aliyunAccessKeySecret"] != "" {
		client, _ = oss.New(aliConfig["aliyunEndPoint"],
			aliConfig["aliyunAccessKeyId"],
			aliConfig["aliyunAccessKeySecret"])
		o.Client = client
		return true
	}
	return false
}

func (o *Ali) WebUploadSign() (*PolicyToken, *errno.Errno) {
	accessKeyId := o.Client.Config.AccessKeyID
	accessKeySecret := o.Client.Config.AccessKeySecret
	ossConfig := make(map[string]interface{}, 0)
	ossConfigStr, _ := model.GetConfigByParamKey("CLOUD_STORAGE_CONFIG_KEY")
	jsoniter.UnmarshalFromString(ossConfigStr.ParamValue, &ossConfig)
	bucketName := (ossConfig["aliyunBucketName"]).(string)
	bucketInfo, _ := o.Client.GetBucketInfo(bucketName)
	host := "http://" + bucketInfo.BucketInfo.Name + "." +
		bucketInfo.BucketInfo.ExtranetEndpoint
	expireTime := int64(30)
	now := time.Now()
	dir := now.Format("20060102") + "/"

	nowTimestamp := now.Unix()
	expireEnd := nowTimestamp + expireTime
	tokenExpire := util.GetGmtIso8601(expireEnd)

	//create post policy json
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, dir)
	pc := PolicyConfig{}
	pc.Expiration = tokenExpire
	pc.Conditions = append(pc.Conditions, condition)
	//calucate signature
	result, err := jsoniter.Marshal(pc)
	if err != nil {
		return nil, errno.ErrOssGenerateSignatureFail
	}
	deByte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash {
		return sha1.New()
	}, []byte(accessKeySecret))
	//io.WriteString(h, debyte)
	h.Write([]byte(deByte))
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	pt := &PolicyToken{
		AccessKeyId: accessKeyId,
		Host:        host,
		Expire:      expireEnd,
		Signature:   signedStr,
		Policy:      deByte,
		Directory:   dir,
	}
	return pt, nil
}
