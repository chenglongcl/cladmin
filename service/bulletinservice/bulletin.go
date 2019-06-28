package bulletinservice

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"cladmin/util"
	"sync"
)

type Bulletin struct {
	Id      uint64
	Title   string
	Tag     string
	Content string
}

func (a *Bulletin) Add() *errno.Errno {
	data := map[string]interface{}{
		"title":   a.Title,
		"tag":     a.Tag,
		"content": a.Content,
	}
	if err := model.AddBulletin(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Bulletin) Edit() *errno.Errno {
	data := map[string]interface{}{
		"id":      a.Id,
		"title":   a.Title,
		"tag":     a.Tag,
		"content": a.Content,
	}
	if err := model.EditBulletin(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Bulletin) Get() (*model.Bulletin, *errno.Errno) {
	publicNotice, err := model.GetBulletin(a.Id)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return publicNotice, nil
}

func (a *Bulletin) GetList(ps util.PageSetting) ([]*model.BulletinInfo, uint64, *errno.Errno) {
	w := make(map[string]interface{})
	if a.Title != "" {
		w["title like"] = "%" + a.Title + "%"
	}
	if a.Tag != "" {
		w["tag"] = a.Tag
	}
	publicNotices, count, err := model.GetBulletinList(w, ps.Offset, ps.Limit)
	if err != nil {
		return nil, count, errno.ErrDatabase
	}
	var ids []uint64
	for _, publicNotice := range publicNotices {
		ids = append(ids, publicNotice.Id)
	}
	wg := sync.WaitGroup{}
	publicNoticesList := model.BulletinList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.BulletinInfo, len(publicNotices)),
	}
	finished := make(chan bool, 1)
	for _, publicNotice := range publicNotices {
		wg.Add(1)
		go func(publicNotice *model.Bulletin) {
			defer wg.Done()
			publicNoticesList.Lock.Lock()
			defer publicNoticesList.Lock.Unlock()
			publicNoticesList.IdMap[publicNotice.Id] = &model.BulletinInfo{
				Id:         publicNotice.Id,
				Title:      publicNotice.Title,
				Tag:        publicNotice.Tag,
				Content:    publicNotice.Content,
				CreateTime: publicNotice.CreatedAt.Format("2006-01-02 15:04:05"),
			}
		}(publicNotice)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}
	info := make([]*model.BulletinInfo, 0)
	for _, id := range ids {
		info = append(info, publicNoticesList.IdMap[id])
	}
	return info, count, nil
}

func (a *Bulletin) Delete() *errno.Errno {
	if err := model.DeleteBulletin(a.Id); err != nil {
		return errno.ErrDatabase
	}
	return nil
}
