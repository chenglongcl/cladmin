package oss

import (
	"cladmin/pkg/oss/client"
	"cladmin/pkg/oss/common"
)

func Init() {
	client.InitAliClient()
}

func SelectClient(name string) common.OSSClient {
	switch name {
	case "ali":
		return client.DefaultAliClient()
	}
	return nil
}
