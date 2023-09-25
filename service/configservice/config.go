package configservice

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
)

type config struct {
	ID             uint64
	ParamKey       string
	ParamValue     string
	Status         int32
	Type           int32
	Remark         string
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Config = *config

func NewConfigService(ctx *gin.Context, opts ...service.Option) Config {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &config{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Config) UpsertByID(updateData map[string]interface{}) (*cladminmodel.SysConfig, *errno.Errno) {
	configModel := &cladminmodel.SysConfig{
		ID:         a.ID,
		ParamKey:   a.ParamKey,
		ParamValue: a.ParamValue,
		Type:       a.Type,
		Status:     a.Status,
		Remark:     a.Remark,
	}
	err := cladminquery.Q.WithContext(a.ctx).SysConfig.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(updateData),
	}).Create(configModel)
	return configModel, gormx.HandleError(err)
}

func (a Config) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysConfig, *errno.Errno) {
	configModel, err := cladminquery.Q.WithContext(a.ctx).SysConfig.Select(fields...).Where(conditions...).Take()
	return configModel, gormx.HandleError(err)
}

func (a Config) Edit(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysConfig.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}
