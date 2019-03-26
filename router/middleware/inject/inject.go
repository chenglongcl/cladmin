package inject

import (
	"cladmin/service/bll"
	"github.com/casbin/casbin"
	"github.com/facebookgo/inject"
	"github.com/lexkong/log"
	"runtime"
)

// Object 注入对象
type Object struct {
	Common   *bll.Common
	Enforcer *casbin.Enforcer
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
	enforcer := casbin.NewEnforcer(path, false)
	_ = g.Provide(&inject.Object{Value: enforcer})
	Common := new(bll.Common)
	_ = g.Provide(&inject.Object{Value: Common})
	if err := g.Populate(); err != nil {
		log.Error("初始化依赖注入发生错误：", err)
	}
	Obj = &Object{
		Enforcer: enforcer,
		Common:   Common,
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
