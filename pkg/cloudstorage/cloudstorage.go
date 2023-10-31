package cloudstorage

import (
	"cladmin/pkg/cloudstorage/aliyunoss"
	"github.com/chenglongcl/log"
)

var cloudStorage *CloudStorage

type CloudStorage struct {
	AliYun *aliyunoss.OSS
}

func InitCloudStorage() {
	var (
		err error
	)
	//
	aliYun := &aliyunoss.OSS{}
	if err = aliYun.NewClient(); err != nil {
		log.Errorf(err, "阿里云OSS客户端初始化失败")
	}
	//
	cloudStorage = &CloudStorage{
		AliYun: aliYun,
	}
}

func GetCloudStorage() *CloudStorage {
	return cloudStorage
}
