package bulletinservice

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

type bulletin struct {
	ID             uint64
	Title          string
	Tag            string
	Content        string
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Bulletin = *bulletin

func NewBulletinService(ctx *gin.Context, opts ...service.Option) Bulletin {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &bulletin{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Bulletin) Add() (*cladminmodel.SysBulletin, *errno.Errno) {
	bulletinModel := &cladminmodel.SysBulletin{ID: 0,
		Title:   a.Title,
		Tag:     a.Tag,
		Content: a.Content,
	}
	err := cladminquery.Q.WithContext(a.ctx).SysBulletin.Create(bulletinModel)
	return bulletinModel, gormx.HandleError(err)
}

func (a Bulletin) Edit(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysBulletin.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}

func (a Bulletin) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysBulletin, *errno.Errno) {
	bulletinModel, err := cladminquery.Q.WithContext(a.ctx).SysBulletin.Select(fields...).Where(conditions...).Take()
	return bulletinModel, gormx.HandleError(err)
}

func (a Bulletin) InfoList(listParams *service.ListParams) ([]*cladminentity.BulletinInfo, uint64, *errno.Errno) {
	bulletinModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, bulletinModel := range bulletinModels {
		ids = append(ids, bulletinModel.ID)
	}
	info := make([]*cladminentity.BulletinInfo, 0)
	wg := sync.WaitGroup{}
	bulletinList := cladminentity.BulletinList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.BulletinInfo, len(bulletinModels)),
	}
	finished := make(chan bool, 1)
	for _, bulletinModel := range bulletinModels {
		wg.Add(1)
		go func(bulletinModel *cladminmodel.SysBulletin) {
			defer wg.Done()
			bulletinList.Lock.Lock()
			defer bulletinList.Lock.Unlock()
			bulletinList.IdMap[bulletinModel.ID] = &cladminentity.BulletinInfo{
				ID:         bulletinModel.ID,
				Title:      bulletinModel.Title,
				Tag:        bulletinModel.Tag,
				Content:    bulletinModel.Content,
				CreateTime: bulletinModel.CreatedAt.Format("2006-01-02 15:04:05"),
			}
		}(bulletinModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		info = append(info, bulletinList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a Bulletin) List(listParams *service.ListParams) (result []*cladminmodel.SysBulletin, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysBulletin
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysBulletin.WithContext(a.ctx)
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

func (a Bulletin) Delete(conditions []gen.Condition) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysBulletin.Where(conditions...).Delete()
	return gormx.HandleError(err)
}
