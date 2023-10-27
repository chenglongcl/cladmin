package client

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/util"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/chenglongcl/log"
	"github.com/json-iterator/go"
	"hash"
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

type OSSConfig struct {
	AliYunAccessKeyID     string `json:"aliyunAccessKeyId"`
	AliYunAccessKeySecret string `json:"aliyunAccessKeySecret"`
	AliYunBucketName      string `json:"aliyunBucketName"`
	AliYunEndPoint        string `json:"aliyunEndPoint"`
}

type Ali struct {
	Client *oss.Client
}

var ali *Ali

func InitAliClient() {
	var (
		client    *oss.Client
		ossConfig OSSConfig
	)
	configModel, err := cladminquery.Q.WithContext(context.Background()).SysConfig.Where(
		cladminquery.Q.SysConfig.ParamKey.Eq("CLOUD_STORAGE_ALI_CONFIG_KEY"),
	).Take()
	if err != nil || configModel == nil || configModel.ID == 0 {
		log.Errorf(err, "获取阿里云OSS配置失败")
		return
	}
	_ = jsoniter.UnmarshalFromString(configModel.ParamValue, &ossConfig)
	client, err = oss.New(ossConfig.AliYunEndPoint, ossConfig.AliYunAccessKeyID, ossConfig.AliYunAccessKeySecret)
	if err != nil {
		log.Errorf(err, "初始化阿里云OSS客户端失败")
		return
	}
	ali = &Ali{
		Client: client,
	}
}

func DefaultAliClient() *Ali {
	return ali
}

func (o *Ali) ResetClient() bool {
	var (
		client    *oss.Client
		ossConfig OSSConfig
	)
	configModel, err := cladminquery.Q.WithContext(context.Background()).SysConfig.Where(
		cladminquery.Q.SysConfig.ParamKey.Eq("CLOUD_STORAGE_ALI_CONFIG_KEY"),
	).Take()
	if err != nil || configModel == nil || configModel.ID == 0 {
		log.Errorf(err, "reset aliYunOSS client error")
		return false
	}
	_ = jsoniter.UnmarshalFromString(configModel.ParamValue, &ossConfig)
	client, err = oss.New(ossConfig.AliYunEndPoint, ossConfig.AliYunAccessKeyID, ossConfig.AliYunAccessKeySecret)
	if err != nil {
		log.Errorf(err, "reset aliYunOSS client error")
		return false
	}
	o.Client = client
	return true
}

func (o *Ali) WebUploadSign() (*PolicyToken, *errno.Errno) {
	accessKeyId := o.Client.Config.AccessKeyID
	accessKeySecret := o.Client.Config.AccessKeySecret
	var ossConfig OSSConfig
	configModel, err := cladminquery.Q.WithContext(context.Background()).SysConfig.Where(
		cladminquery.Q.SysConfig.ParamKey.Eq("CLOUD_STORAGE_ALI_CONFIG_KEY"),
	).Take()
	if err != nil || configModel == nil || configModel.ID == 0 {
		return nil, errno.ErrOssGenerateSignatureFail
	}
	_ = jsoniter.UnmarshalFromString(configModel.ParamValue, &ossConfig)
	bucketName := ossConfig.AliYunBucketName
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
	//calculate signature
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
