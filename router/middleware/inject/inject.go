package inject

import (
	"cladmin/service/bll"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/casbin/casbin"
	"github.com/facebookgo/inject"
	"github.com/lexkong/log"
	"runtime"
)

// Object 注入对象
type Object struct {
	Common          *bll.Common
	Enforcer        *casbin.Enforcer
	AliYunOssClient *oss.Client
}

var Obj *Object

// Init 初始化依赖注入
func Init() {
	g := new(inject.Graph)
	// 注入casbin
	osType := runtime.GOOS
	var path string
	if osType == "windows" {
		path = "conf\\rbac_model.conf"
	} else if osType == "linux" {
		path = "conf/rbac_model.conf"
	}
	//casbin new
	enforcer := casbin.NewEnforcer(path, false)
	_ = g.Provide(&inject.Object{Value: enforcer})
	//aliyun oss new
	aliYunOssClient, _ := oss.New("http://oss-cn-shanghai.aliyuncs.com",
		"LTAIcpkAxHSr8L6t",
		"tI2UWkXUzutGMtVeIWSzCrX5IKmONS")
	_ = g.Provide(&inject.Object{Value: aliYunOssClient})
	//common new
	Common := new(bll.Common)
	_ = g.Provide(&inject.Object{Value: Common})
	if err := g.Populate(); err != nil {
		log.Error("初始化依赖注入发生错误：", err)
	}
	Obj = &Object{
		Common:          Common,
		Enforcer:        enforcer,
		AliYunOssClient: aliYunOssClient,
	}
	return
}

// 加载casbin策略数据，包括角色权限数据、用户角色数据
func LoadCasbinPolicyData() error {
	c := Obj.Common
	err := c.RoleAPI.LoadAllPolicy()
	if err != nil {
		return err
	}
	err = c.UserAPI.LoadAllPolicy()
	if err != nil {
		return err
	}
	return nil
}
