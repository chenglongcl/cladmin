package articleservice

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
	"time"
)

type article struct {
	ID             uint64
	UserID         uint64
	CateID         uint64
	Title          string
	Thumb          string
	Content        string
	ReleaseTime    string
	serviceOptions *service.Options
	ctx            *gin.Context
}

type Article = *article

func NewArticleService(ctx *gin.Context, opts ...service.Option) Article {
	opt := new(service.Options)
	for _, f := range opts {
		f(opt)
	}
	return &article{
		serviceOptions: opt,
		ctx:            ctx,
	}
}

func (a Article) Add() (*cladminmodel.SysArticle, *errno.Errno) {
	articleModel := &cladminmodel.SysArticle{
		UserID:      a.UserID,
		CateID:      a.CateID,
		Title:       a.Title,
		Thumb:       a.Thumb,
		Content:     a.Content,
		ReleaseTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := cladminquery.Q.WithContext(a.ctx).SysArticle.Create(articleModel)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, errNo
	}
	return articleModel, nil
}

func (a Article) Edit(conditions []gen.Condition, data map[string]interface{}) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysArticle.Where(conditions...).Updates(data)
	return gormx.HandleError(err)
}

func (a Article) Get(fields []field.Expr, conditions []gen.Condition) (*cladminmodel.SysArticle, *errno.Errno) {
	articleModel, err := cladminquery.Q.WithContext(a.ctx).SysArticle.Select(fields...).Where(conditions...).Take()
	return articleModel, gormx.HandleError(err)
}

func (a Article) InfoList(listParams *service.ListParams) ([]*cladminentity.ArticleInfo, uint64, *errno.Errno) {
	articleModels, count, err := a.List(listParams)
	if errNo := gormx.HandleError(err); errNo != nil {
		return nil, uint64(count), errNo
	}
	var ids []uint64
	for _, articleModel := range articleModels {
		ids = append(ids, articleModel.ID)
	}
	info := make([]*cladminentity.ArticleInfo, 0)
	wg := sync.WaitGroup{}
	articleList := cladminentity.ArticleList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*cladminentity.ArticleInfo, len(articleModels)),
	}
	finished := make(chan bool, 1)
	for _, articleModel := range articleModels {
		wg.Add(1)
		go func(articleModel *cladminmodel.SysArticle) {
			defer wg.Done()
			articleList.Lock.Lock()
			defer articleList.Lock.Unlock()
			articleList.IdMap[articleModel.ID] = &cladminentity.ArticleInfo{
				ID:          articleModel.ID,
				UserID:      articleModel.UserID,
				CateID:      articleModel.CateID,
				Title:       articleModel.Title,
				Thumb:       articleModel.Thumb,
				ReleaseTime: articleModel.ReleaseTime,
			}
		}(articleModel)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		info = append(info, articleList.IdMap[id])
	}
	return info, uint64(count), nil
}

func (a Article) List(listParams *service.ListParams) (result []*cladminmodel.SysArticle, count int64, err error) {
	qc := cladminquery.Q.WithContext(a.ctx).SysArticle
	if listParams.Options.CustomDBOrder != "" {
		qc = cladminquery.Q.SysArticle.WithContext(a.ctx)
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

func (a Article) Delete(conditions []gen.Condition) *errno.Errno {
	_, err := cladminquery.Q.WithContext(a.ctx).SysArticle.Where(conditions...).Delete()
	return gormx.HandleError(err)
}
