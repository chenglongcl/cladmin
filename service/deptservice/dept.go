package deptservice

import (
	"cladmin/dal/cladmindb/cladminentity"
	"cladmin/dal/cladmindb/cladminmodel"
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/pkg/errno"
	"cladmin/pkg/gormx"
	"cladmin/pkg/tree"
	"cladmin/service"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"sync"
)

type dept struct {
	ID             uint64
	ParentID       uint64
	Name           string
	Sort           int32
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Dept = *dept

func NewDeptService(ctx *gin.Context, opts ...service.Option) Dept {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &dept{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Dept) Add() (*cladminmodel.SysDept, *errno.Errno) {
	deptModel := &cladminmodel.SysDept{
		ParentID: a.ParentID,
		Name:     a.Name,
		Sort:     a.Sort,
	}
	err := cladminquery.Q.WithContext(a.ctx).SysDept.Create(deptModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	return deptModel, nil
}

func (a Dept) Edit(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysDept.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}

func (a Dept) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysDept, *errno.Errno) {
	deptModel, err := cladminquery.Q.WithContext(a.ctx).SysDept.Select(fields...).Where(conditions...).Take()
	return deptModel, gormx.HandleError(err)
}

func (a Dept) InfoList(listParams *service.ListParams) ([]*cladminentity.DeptInfo, uint64, *errno.Errno) {
	deptModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, deptModel := range deptModels {
		ids = append(ids, deptModel.ID)
	}
	info := make([]*cladminentity.DeptInfo, 0)
	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	deptList := cladminentity.DeptList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.DeptInfo, len(deptModels)),
	}
	for _, deptModel := range deptModels {
		wg.Add(1)
		go func(deptModel *cladminmodel.SysDept) {
			defer wg.Done()
			deptList.Lock.Lock()
			defer deptList.Lock.Unlock()
			deptList.IdMap[deptModel.ID] = &cladminentity.DeptInfo{
				ID:       deptModel.ID,
				ParentID: deptModel.ParentID,
				Name:     deptModel.Name,
				Sore:     deptModel.Sort,
			}
		}(deptModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}

	for _, id := range ids {
		info = append(info, deptList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a Dept) Tree(listParams *service.ListParams) ([]*cladminentity.DeptTree, *errno.Errno) {
	deptModels, _, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	data := make([]*cladminentity.DeptTree, 0)
	for _, deptModel := range deptModels {
		data = append(data, &cladminentity.DeptTree{
			ID:       deptModel.ID,
			ParentID: deptModel.ParentID,
			Name:     deptModel.Name,
			Sort:     deptModel.Sort,
		})
	}
	return tree.ToTree[cladminentity.DeptTree](data, &cladminentity.DeptTree{}), nil
}

func (a Dept) List(listParams *service.ListParams) (result []*cladminmodel.SysDept, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysDept
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysDept.WithContext(a.ctx)
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

func (a Dept) Delete(conditions []gen.Condition) *errno.Errno {
	// 查询部门是否存在
	deptModel, err := cladminquery.Q.WithContext(a.ctx).SysDept.Where(conditions...).Take()
	if err != nil {
		return gormx.HandleError(err)
	}
	if deptModel == nil || deptModel.ID == 0 {
		return errno.ErrRecordNotFound
	}
	//删除部门
	_, err = cladminquery.Q.WithContext(a.ctx).SysDept.Where(conditions...).Unscoped().Delete()
	if err != nil {
		return gormx.HandleError(err)
	}
	//更新部门下的用户
	_, _ = cladminquery.Q.WithContext(a.ctx).SysUser.Where(
		cladminquery.Q.SysUser.DeptID.Eq(deptModel.ID),
	).UpdateSimple(cladminquery.Q.SysUser.DeptID.Zero())
	return nil
}

func (a Dept) Count(conditions []gen.Condition) (int64, *errno.Errno) {
	count, err := cladminquery.Q.WithContext(a.ctx).SysDept.Where(conditions...).Count()
	return count, gormx.HandleError(err)
}
