package article_service

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"cladmin/util"
	"sync"
	"time"
)

type Article struct {
	Id          uint64
	UserId      uint64
	CateId      uint64
	Title       string
	Thumb       string
	Content     string
	ReleaseTime string
}

func (a *Article) Add() *errno.Errno {
	data := map[string]interface{}{
		"user_id":      a.UserId,
		"cate_id":      a.CateId,
		"title":        a.Title,
		"thumb":        a.Thumb,
		"content":      a.Content,
		"release_time": time.Now().Format("2006-01-02 15:03:04"),
	}
	if err := model.AddArticle(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Article) Edit() *errno.Errno {
	data := map[string]interface{}{
		"id":           a.Id,
		"user_id":      a.UserId,
		"cate_id":      a.CateId,
		"title":        a.Title,
		"thumb":        a.Thumb,
		"content":      a.Content,
		"release_time": a.ReleaseTime,
	}
	if err := model.EditArticle(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Article) Get() (*model.Article, *errno.Errno) {
	article, err := model.GetArticle(a.Id)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return article, nil
}

func (a *Article) GetList(ps util.PageSetting) ([]*model.ArticleInfo, uint64, *errno.Errno) {
	w := make(map[string]interface{})
	if a.Title != "" {
		w["title like"] = "%" + a.Title + "%"
	}
	if a.CateId != 0 {
		w["cate_id"] = a.CateId
	}
	articles, count, err := model.GetArticleList(w, ps.Offset, ps.Limit)
	if err != nil {
		return nil, count, errno.ErrDatabase
	}
	var ids []uint64
	for _, article := range articles {
		ids = append(ids, article.Id)
	}

	info := make([]*model.ArticleInfo, 0)
	wg := sync.WaitGroup{}
	ArticleList := model.ArticleList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.ArticleInfo, len(articles)),
	}
	finished := make(chan bool, 1)

	for _, article := range articles {
		wg.Add(1)
		go func(article *model.Article) {
			defer wg.Done()
			ArticleList.Lock.Lock()
			defer ArticleList.Lock.Unlock()
			ArticleList.IdMap[article.Id] = &model.ArticleInfo{
				Id:          article.Id,
				UserId:      article.UserId,
				CateId:      article.CateId,
				Title:       article.Title,
				Thumb:       article.Thumb,
				ReleaseTime: article.ReleaseTime,
			}
		}(article)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		info = append(info, ArticleList.IdMap[id])
	}
	return info, count, nil
}

func (a *Article) Delete() *errno.Errno {
	if err := model.DeleteArticle(a.Id); err != nil {
		return errno.ErrDatabase
	}
	return nil
}
