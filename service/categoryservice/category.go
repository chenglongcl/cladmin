package categoryservice

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

type category struct {
	ID             uint64
	ParentID       uint64
	Name           string
	Icon           string
	OrderNum       int64
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Category = *category

func NewCategoryService(ctx *gin.Context, opts ...service.Option) Category {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &category{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Category) Add() (*cladminmodel.SysCategory, *errno.Errno) {
	categoryModel := &cladminmodel.SysCategory{
		ParentID: a.ParentID,
		Name:     a.Name,
		Icon:     a.Icon,
		OrderNum: a.OrderNum,
	}
	err := cladminquery.Q.WithContext(a.ctx).SysCategory.Create(categoryModel)
	return categoryModel, gormx.HandleError(err)
}

func (a Category) Edit(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysCategory.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}

func (a Category) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysCategory, *errno.Errno) {
	categoryModel, err := cladminquery.Q.WithContext(a.ctx).SysCategory.Select(fields...).Where(conditions...).Take()
	return categoryModel, gormx.HandleError(err)
}

func (a Category) InfoList(listParams *service.ListParams) ([]*cladminentity.CategoryInfo, uint64, *errno.Errno) {
	categoryModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, categoryModel := range categoryModels {
		ids = append(ids, categoryModel.ID)
	}
	info := make([]*cladminentity.CategoryInfo, 0)
	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	categoryList := cladminentity.CategoryList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.CategoryInfo, len(categoryModels)),
	}
	for _, categoryModel := range categoryModels {
		wg.Add(1)
		go func(categoryModel *cladminmodel.SysCategory) {
			defer wg.Done()
			categoryList.Lock.Lock()
			defer categoryList.Lock.Unlock()
			categoryList.IdMap[categoryModel.ID] = &cladminentity.CategoryInfo{
				Id:         categoryModel.ID,
				ParentID:   categoryModel.ParentID,
				Name:       categoryModel.Name,
				Icon:       categoryModel.Icon,
				OrderNum:   categoryModel.OrderNum,
				CreateTime: categoryModel.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdateTime: categoryModel.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(categoryModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		info = append(info, categoryList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a Category) List(listParams *service.ListParams) (result []*cladminmodel.SysCategory, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysCategory
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysCategory.WithContext(a.ctx)
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

func (a Category) Delete(categoryModel *cladminmodel.SysCategory) *errno.Errno {
	subCategoryModel, errNo := a.Get([]field.Expr{
		cladminquery.Q.SysCategory.ID,
	}, []gen.Condition{
		cladminquery.Q.SysCategory.ParentID.Eq(categoryModel.ID),
	})
	if errNo != nil {
		return errNo
	}
	if subCategoryModel != nil && subCategoryModel.ID != 0 {
		return errno.ErrRecordHasChildren
	}
	_, err := cladminquery.Q.WithContext(a.ctx).SysCategory.Delete(categoryModel)
	return gormx.HandleError(err)
}
