package public_notice_service

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"cladmin/util"
	"sync"
)

type PublicNotice struct {
	Id      uint64
	Title   string
	Tag     string
	Content string
}

func (a *PublicNotice) Add() *errno.Errno {
	data := map[string]interface{}{
		"title":   a.Title,
		"tag":     a.Tag,
		"content": a.Content,
	}
	if err := model.AddPublicNotice(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *PublicNotice) Edit() *errno.Errno {
	data := map[string]interface{}{
		"id":      a.Id,
		"title":   a.Title,
		"tag":     a.Tag,
		"content": a.Content,
	}
	if err := model.EditPublicNotice(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *PublicNotice) Get() (*model.PublicNotice, *errno.Errno) {
	publicNotice, err := model.GetPublicNotice(a.Id)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return publicNotice, nil
}

func (a *PublicNotice) GetList(ps util.PageSetting) ([]*model.PublicNoticeInfo, uint64, *errno.Errno) {
	w := make(map[string]interface{})
	if a.Title != "" {
		w["title like"] = "%" + a.Title + "%"
	}
	if a.Tag != "" {
		w["tag"] = a.Tag
	}
	publicNotices, count, err := model.GetPublicNoticeList(w, ps.Offset, ps.Limit)
	if err != nil {
		return nil, count, errno.ErrDatabase
	}
	var ids []uint64
	for _, publicNotice := range publicNotices {
		ids = append(ids, publicNotice.Id)
	}
	wg := sync.WaitGroup{}
	publicNoticesList := model.PublicNoticeList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.PublicNoticeInfo, len(publicNotices)),
	}
	finished := make(chan bool, 1)
	for _, publicNotice := range publicNotices {
		wg.Add(1)
		go func(publicNotice *model.PublicNotice) {
			defer wg.Done()
			publicNoticesList.Lock.Lock()
			defer publicNoticesList.Lock.Unlock()
			publicNoticesList.IdMap[publicNotice.Id] = &model.PublicNoticeInfo{
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
	info := make([]*model.PublicNoticeInfo, 0)
	for _, id := range ids {
		info = append(info, publicNoticesList.IdMap[id])
	}
	return info, count, nil
}

func (a *PublicNotice) Delete() *errno.Errno {
	if err := model.DeletePublicNotice(a.Id); err != nil {
		return errno.ErrDatabase
	}
	return nil
}
