package demoservice

import (
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/pkg/redisgo"
	"cladmin/service"
	"cladmin/service/userservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type demo struct {
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Demo = *demo

func NewDemoService(ctx *gin.Context, opts ...service.Option) Demo {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &demo{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Demo) DemoOne() (*cladminmodel.SysUser, *errno.Errno) {
	userService := userservice.NewUserService(a.ctx)
	userModel, errNo := userService.Get([]field.Expr{
		cladminquery.Q.SysUser.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysUser.ID.Eq(1),
	})
	if errNo != nil {
		return nil, errNo
	}
	_, _ = redisgo.My().HSet("testUsers", "1", userModel)
	userModelTwo := &cladminmodel.SysUser{}
	_ = redisgo.My().HGetObject("testUsers", "1", userModelTwo)
	return userModelTwo, nil
}
