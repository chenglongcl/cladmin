package service

import (
	"apiserver/util"
	"apiserver/model"
	"sync"
	"github.com/json-iterator/go"
)

func ListArticle(cateId uint64, ps util.PageSetting) ([]*model.ArticleInfo, uint64, error) {
	infos := make([]*model.ArticleInfo, 0)
	articles, count, err := model.ListArticle(cateId, ps.Offset, ps.Limit)
	if err != nil {
		return nil, count, err
	}
	var ids []uint64
	for _, article := range articles {
		ids = append(ids, article.Id)
	}

	wg := sync.WaitGroup{}
	articleList := model.ArticleList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.ArticleInfo, len(articles)),
	}
	finished := make(chan bool, 1)

	for _, a := range articles {
		wg.Add(1)
		go func(a *model.ArticleModel) {
			defer wg.Done()
			articleList.Lock.Lock()
			defer articleList.Lock.Unlock()
			articleList.IdMap[a.Id] = &model.ArticleInfo{
				Id:        a.Id,
				CateId:    a.CateId,
				Title:     a.Title,
				Images:    jsonParse(a),
				Author:    model.Author{Id: a.Author.Id, Username: a.Author.Username},
				CreatedAt: a.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: a.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(a)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	for _, id := range ids {
		infos = append(infos, articleList.IdMap[id])
	}
	return infos, count, nil
}

func jsonParse(a *model.ArticleModel) (images []string) {
	jsoniter.Unmarshal([]byte(a.Images), &images)
	return
}
