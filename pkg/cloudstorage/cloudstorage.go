package cloudstorage

import (
	"cladmin/pkg/cloudstorage/client"
	"cladmin/pkg/cloudstorage/iface"
)

func Init() {
	client.InitAliClient()
}

func SelectClient(name string) iface.IClient {
	switch name {
	case "ali":
		return client.DefaultAliClient()
	}
	return nil
}
