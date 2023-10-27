package dictservice

import (
	"cladmin/dal/cladmindb/cladminentity"
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/pkg/gormx"
	"cladmin/service"
	"cladmin/util"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"sync"
)

type dictType struct {
	ID             uint64
	DictType       string
	DictName       string
	Remark         string
	Sort           uint64
	serviceOptions *service.Options
	ctx            *gin.Context
}

type DictType = *dictType

func NewDictTypeService(ctx *gin.Context, opts ...service.Option) DictType {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &dictType{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a DictType) Add() (*cladminmodel.SysDictType, *errno.Errno) {
	dictTypeModel := &cladminmodel.SysDictType{
		ID:       a.ID,
		DictType: a.DictType,
		DictName: a.DictName,
		Remark:   a.Remark,
		Sort:     a.Sort,
	}
	err := cladminquery.Q.WithContext(a.ctx).SysDictType.Create(dictTypeModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	return dictTypeModel, nil
}

func (a DictType) Edit(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysDictType.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}

func (a DictType) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysDictType, *errno.Errno) {
	dictTypeModel, err := cladminquery.Q.WithContext(a.ctx).SysDictType.Select(fields...).Where(conditions...).Take()
	return dictTypeModel, gormx.HandleError(err)
}

func (a DictType) InfoList(listParams *service.ListParams) ([]*cladminentity.DictTypeInfo, uint64, *errno.Errno) {
	dictTypeModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var (
		ids []uint64
	)
	for _, dictTypeModel := range dictTypeModels {
		ids = append(ids, dictTypeModel.ID)
	}
	dictDataInfos := make([]*cladminentity.DictDataInfo, 0)
	if listParams.Options.Scenes == "all" {
		dictDataService := NewDictDataService(a.ctx)
		dictDataInfos, _, _ = dictDataService.InfoList(&service.ListParams{
			PS: util.PageSetting{},
			Options: struct {
				WithoutCount  bool
				Scenes        string
				CustomDBOrder string
				CustomFunc    func() interface{}
			}{WithoutCount: true},
			Fields: []field.Expr{cladminquery.Q.SysDictData.ALL},
			Conditions: []gen.Condition{
				cladminquery.Q.SysDictData.DictTypeID.In(ids...),
			},
			Orders: []field.Expr{
				cladminquery.Q.SysDictData.Sort,
			},
		})
	}
	info := make([]*cladminentity.DictTypeInfo, 0)
	wg := sync.WaitGroup{}
	dictTypeList := cladminentity.DictTypeList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.DictTypeInfo, len(dictTypeModels)),
	}
	finished := make(chan bool, 1)
	for _, dictTypeModel := range dictTypeModels {
		wg.Add(1)
		go func(dictTypeModel *cladminmodel.SysDictType) {
			defer wg.Done()
			dictTypeList.Lock.Lock()
			defer dictTypeList.Lock.Unlock()
			dictTypeList.IdMap[dictTypeModel.ID] = &cladminentity.DictTypeInfo{
				ID:         dictTypeModel.ID,
				DictType:   dictTypeModel.DictType,
				DictName:   dictTypeModel.DictName,
				Remark:     dictTypeModel.Remark,
				Sort:       dictTypeModel.Sort,
				CreateTime: dictTypeModel.CreatedAt.Format("2006-01-02 15:04:05"),
				DataList: slice.Filter(dictDataInfos, func(i int, dictDataInfo *cladminentity.DictDataInfo) bool {
					return dictDataInfo.DictTypeID == dictTypeModel.ID
				}),
			}

		}(dictTypeModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		info = append(info, dictTypeList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a DictType) List(listParams *service.ListParams) (result []*cladminmodel.SysDictType, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysDictType
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysDictType.WithContext(a.ctx)
		qc.ReplaceDB(qc.UnderlyingDB().Order(listParams.Options.CustomDBOrder))
	}
	base := qc.Select(listParams.Fields...).Where(listParams.Conditions...).Order(listParams.Orders...)
	for _, leftJoin := range listParams.LeftJoins {
		base = base.LeftJoin(leftJoin.Table, leftJoin.On...)
	}
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

func (a DictType) Delete(conditions []gen.Condition) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysDictType.Where(conditions...).Delete()
	return gormx.HandleError(err)
}
