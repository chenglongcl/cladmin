package inject

import (
	"cladmin/service/bll"
	"github.com/casbin/casbin"
	"github.com/chenglongcl/log"
	"github.com/facebookgo/inject"
	"runtime"
)

// Object
// @Description: 注入对象
type Object struct {
	Common   *bll.Common
	Enforcer *casbin.Enforcer
}

var Obj *Object

// Init
// @Description: 初始化依赖注入
func Init() {
	defer func() {
		if err := recover(); err != nil {
			log.Infof("%+v", err)
		}
	}()
	var g inject.Graph
	// 注入casbin
	osType := runtime.GOOS
	var path string
	switch osType {
	case "windows":
		path = "conf\\rbac_model.conf"
	case "linux":
		path = "conf/rbac_model.conf"
	case "darwin":
		path = "conf/rbac_model.conf"
	}
	//casbin new
	enforcer := casbin.NewEnforcer(path, false)
	_ = g.Provide(&inject.Object{Value: enforcer})
	//common new
	common := new(bll.Common)
	_ = g.Provide(&inject.Object{Value: common})
	if err := g.Populate(); err != nil {
		log.Error("初始化依赖注入发生错误：", err)
	}
	Obj = &Object{
		Common:   common,
		Enforcer: enforcer,
	}
	return
}

// LoadCasbinPolicyData
// @Description: 加载casbin策略数据，包括角色权限数据、用户角色数据
// @return error
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
