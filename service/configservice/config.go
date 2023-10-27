package configservice

import (
	"cladmin/dal/cladmindb/cladminentity"
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/pkg/gormx"
	"cladmin/service"
	"cladmin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
	"sync"
)

type config struct {
	ID             uint64
	ParamKey       string
	ParamValue     string
	Type           int32
	Status         bool
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

func (a Config) Add() (*cladminmodel.SysConfig, *errno.Errno) {
	configModel := &cladminmodel.SysConfig{
		ParamKey:   a.ParamKey,
		ParamValue: a.ParamValue,
		Type:       a.Type,
		Status:     a.Status,
		Remark:     a.Remark,
	}
	fmt.Printf("%+v\n", configModel)
	err := cladminquery.Q.WithContext(a.ctx).SysConfig.Create(configModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	return configModel, nil
}

func (a Config) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysConfig, *errno.Errno) {
	configModel, err := cladminquery.Q.WithContext(a.ctx).SysConfig.Select(fields...).Where(conditions...).Take()
	return configModel, gormx.HandleError(err)
}

func (a Config) Edit(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysConfig.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}

func (a Config) InfoList(listParams *service.ListParams) ([]*cladminentity.ConfigInfo, uint64, *errno.Errno) {
	configModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, configModel := range configModels {
		ids = append(ids, configModel.ID)
	}
	info := make([]*cladminentity.ConfigInfo, 0)
	wg := sync.WaitGroup{}
	configList := cladminentity.ConfigList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.ConfigInfo, len(configModels)),
	}
	finished := make(chan bool, 1)
	for _, configModel := range configModels {
		wg.Add(1)
		go func(configModel *cladminmodel.SysConfig) {
			defer wg.Done()
			configList.Lock.Lock()
			defer configList.Lock.Unlock()
			configList.IdMap[configModel.ID] = &cladminentity.ConfigInfo{
				ID:         configModel.ID,
				ParamKey:   configModel.ParamKey,
				ParamValue: configModel.ParamValue,
				Type:       configModel.Type,
				Status:     configModel.Status,
				Remark:     configModel.Remark,
				Locked:     configModel.Locked,
			}
		}(configModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		info = append(info, configList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a Config) List(listParams *service.ListParams) (result []*cladminmodel.SysConfig, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysConfig
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysConfig.WithContext(a.ctx)
		qc.ReplaceDB(qc.UnderlyingDB().Order(listParams.Options.CustomDBOrder))
	}
	base := qc.Select(listParams.Fields...).Where(listParams.Conditions...).Order(listParams.Orders...)
	offset, limit := util.MysqlPagination(listParams.PS)
	if !listParams.Options.WithoutCount {
		result, count, err = base.FindByPage(offset, limit)
	} else {
		if limit == -1 {
			result, err = base.Find()
		} else {
			result, err = base.Offset(offset).Limit(limit).Find()
		}
	}
	return
}

func (a Config) Delete(conditions []gen.Condition) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysConfig.Where(conditions...).Delete()
	return gormx.HandleError(err)
}
