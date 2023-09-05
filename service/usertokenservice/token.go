package usertokenservice

import (
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/pkg/gormx"
	"cladmin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
	"time"
)

type userToken struct {
	UserID         uint64
	Token          string
	ExpireTime     time.Time
	RefreshTime    time.Time
	serviceOptions *service.Options
	ctx            *gin.Context
}

type UserToken = *userToken

func NewUserTokenService(ctx *gin.Context, opts ...service.Option) UserToken {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &userToken{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a UserToken) RecordToken() (*cladminmodel.SysUserToken, *errno.Errno) {
	userTokenModel := &cladminmodel.SysUserToken{
		UserID:      a.UserID,
		Token:       a.Token,
		ExpireTime:  a.ExpireTime,
		RefreshTime: a.RefreshTime,
	}
	err := cladminquery.Q.WithContext(a.ctx).SysUserToken.
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"token":        a.Token,
				"expire_time":  a.ExpireTime,
				"refresh_time": a.RefreshTime,
			}),
		}).Create(userTokenModel)
	return userTokenModel, gormx.HandleError(err)
}

func (a UserToken) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysUserToken, *errno.Errno) {
	userTokenModel, err := cladminquery.Q.WithContext(a.ctx).SysUserToken.Select(fields...).Where(conditions...).Take()
	return userTokenModel, gormx.HandleError(err)
}
