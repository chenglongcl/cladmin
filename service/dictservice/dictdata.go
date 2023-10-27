package dictservice

import (
	"cladmin/dal/cladmindb/cladminentity"
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/pkg/gormx"
	"cladmin/service"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"sync"
)

type dictData struct {
	ID             uint64
	DictTypeID     uint64
	DictLabel      string
	DictValue      string
	Remark         string
	Sort           uint64
	serviceOptions *service.Options
	ctx            *gin.Context
}

type DictData = *dictData

func NewDictDataService(ctx *gin.Context, opts ...service.Option) DictData {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &dictData{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a DictData) Add() (*cladminmodel.SysDictData, *errno.Errno) {
	dictDataModel := &cladminmodel.SysDictData{
		ID:         a.ID,
		DictTypeID: a.DictTypeID,
		DictLabel:  a.DictLabel,
		DictValue:  a.DictValue,
		Remark:     a.Remark,
		Sort:       a.Sort,
	}
	err := cladminquery.Q.WithContext(a.ctx).SysDictData.Create(dictDataModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	return dictDataModel, nil
}

func (a DictData) Edit(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysDictData.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}

func (a DictData) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysDictData, *errno.Errno) {
	dictDataModel, err := cladminquery.Q.WithContext(a.ctx).SysDictData.Select(fields...).Where(conditions...).Take()
	return dictDataModel, gormx.HandleError(err)
}

func (a DictData) InfoList(listParams *service.ListParams) ([]*cladminentity.DictDataInfo, uint64, *errno.Errno) {
	dictDataModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, dictDataModel := range dictDataModels {
		ids = append(ids, dictDataModel.ID)
	}
	info := make([]*cladminentity.DictDataInfo, 0)
	wg := sync.WaitGroup{}
	dictDataList := cladminentity.DictDataList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.DictDataInfo, len(dictDataModels)),
	}
	finished := make(chan bool, 1)
	for _, dictDataModel := range dictDataModels {
		wg.Add(1)
		go func(dictDataModel *cladminmodel.SysDictData) {
			defer wg.Done()
			dictDataList.Lock.Lock()
			defer dictDataList.Lock.Unlock()
			dictDataList.IdMap[dictDataModel.ID] = &cladminentity.DictDataInfo{
				ID:         dictDataModel.ID,
				DictTypeID: dictDataModel.DictTypeID,
				DictLabel:  dictDataModel.DictLabel,
				DictValue:  dictDataModel.DictValue,
				Remark:     dictDataModel.Remark,
				Sort:       dictDataModel.Sort,
				CreateTime: dictDataModel.CreatedAt.Format("2006-01-02 15:04:05"),
			}
		}(dictDataModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		info = append(info, dictDataList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a DictData) List(listParams *service.ListParams) (result []*cladminmodel.SysDictData, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysDictData
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysDictData.WithContext(a.ctx)
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

func (a DictData) Delete(conditions []gen.Condition) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysDictData.Where(conditions...).Delete()
	return gormx.HandleError(err)
}
