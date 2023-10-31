package aliyunoss

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/util"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/chenglongcl/log"
	jsoniter "github.com/json-iterator/go"
	"hash"
	"net/http"
	"net/url"
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

type Config struct {
	AliYunAccessKeyID     string `json:"aliyunAccessKeyId"`
	AliYunAccessKeySecret string `json:"aliyunAccessKeySecret"`
	AliYunBucketName      string `json:"aliyunBucketName"`
	AliYunEndPoint        string `json:"aliyunEndPoint"`
}

type OSS struct {
	Client *oss.Client
}

func (a *OSS) NewClient() error {
	var (
		client    *oss.Client
		ossConfig Config
	)
	configModel, err := cladminquery.Q.WithContext(context.Background()).SysConfig.Where(
		cladminquery.Q.SysConfig.ParamKey.Eq("CLOUD_STORAGE_ALI_CONFIG_KEY"),
	).Take()
	if err != nil {
		return err
	}
	if configModel == nil || configModel.ID == 0 {
		return errors.New("阿里云OSS配置不存在")
	}
	_ = jsoniter.UnmarshalFromString(configModel.ParamValue, &ossConfig)
	if client, err = oss.New(
		ossConfig.AliYunEndPoint,
		ossConfig.AliYunAccessKeyID,
		ossConfig.AliYunAccessKeySecret,
	); err != nil {
		return err
	}
	a.Client = client
	return nil
}

func (a *OSS) ResetClient() error {
	if err := a.NewClient(); err != nil {
		log.Errorf(err, "阿里云OSS客户端重置失败")
		return err
	}
	return nil
}

func (a *OSS) WebUploadSign() (*PolicyToken, *errno.Errno) {
	accessKeyId := a.Client.Config.AccessKeyID
	accessKeySecret := a.Client.Config.AccessKeySecret
	var ossConfig Config
	configModel, err := cladminquery.Q.WithContext(context.Background()).SysConfig.Where(
		cladminquery.Q.SysConfig.ParamKey.Eq("CLOUD_STORAGE_ALI_CONFIG_KEY"),
	).Take()
	if err != nil || configModel == nil || configModel.ID == 0 {
		return nil, errno.ErrOssGenerateSignatureFail
	}
	_ = jsoniter.UnmarshalFromString(configModel.ParamValue, &ossConfig)
	bucketName := ossConfig.AliYunBucketName
	bucketInfo, _ := a.Client.GetBucketInfo(bucketName)
	host := "https://" + bucketInfo.BucketInfo.Name + "." +
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

func (a *OSS) IsObjectExist(objectURL string) (bool, error) {
	u, err := url.Parse(objectURL)
	if err != nil {
		return false, err
	}
	bucketName := strings.Split(u.Hostname(), ".")[0]
	objectName := strings.TrimLeft(u.Path, "/")
	bucket, err := a.Client.Bucket(bucketName)
	if err != nil {
		return false, err
	}
	return bucket.IsObjectExist(objectName)
}

func (a *OSS) GetObjectDetailedMeta(objectURL string) (http.Header, error) {
	u, err := url.Parse(objectURL)
	if err != nil {
		return nil, err
	}
	bucketName := strings.Split(u.Hostname(), ".")[0]
	objectName := strings.TrimLeft(u.Path, "/")
	bucket, err := a.Client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket.GetObjectDetailedMeta(objectName)
}
