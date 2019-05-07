package category_service

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"sync"
)

type Category struct {
	Id       uint64
	ParentId uint64
	Name     string
	Icon     string
	OrderNum int64
}

func (a *Category) Add() *errno.Errno {
	data := map[string]interface{}{
		"parent_id": a.ParentId,
		"name":      a.Name,
		"icon":      a.Icon,
		"order_num": a.OrderNum,
	}
	if err := model.AddCategory(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Category) Edit() *errno.Errno {
	data := map[string]interface{}{
		"id":        a.Id,
		"parent_id": a.ParentId,
		"name":      a.Name,
		"icon":      a.Icon,
		"order_num": a.OrderNum,
	}
	if err := model.EditCategory(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Category) Get() (*model.Category, *errno.Errno) {
	category, err := model.GetCategory(a.Id)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return category, nil
}

func (a *Category) GetList(w map[string]interface{}) ([]*model.CategoryInfo, *errno.Errno) {
	categories, err := model.GetCategoryList(w)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	var ids []uint64
	for _, category := range categories {
		ids = append(ids, category.Id)
	}

	info := make([]*model.CategoryInfo, 0)
	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	categoryList := model.CategoryList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.CategoryInfo, len(categories)),
	}

	for _, category := range categories {
		wg.Add(1)
		go func(category *model.Category) {
			defer wg.Done()
			categoryList.Lock.Lock()
			defer categoryList.Lock.Unlock()
			categoryList.IdMap[category.Id] = &model.CategoryInfo{
				Id:         category.Id,
				ParentId:   category.ParentId,
				Name:       category.Name,
				Icon:       category.Icon,
				OrderNum:   category.OrderNum,
				CreateTime: category.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdateTime: category.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(category)
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
	return info, nil
}

func (a *Category) Delete() (*errno.Errno) {
	children, err := model.GetCategoryList(map[string]interface{}{
		"parent_id": a.Id,
	})
	if err != nil {
		return errno.ErrDatabase
	}
	if len(children) > 0 {
		return errno.ErrRecordHasChildren
	}
	if err := model.DeleteCategory(a.Id); err != nil {
		return errno.ErrDatabase
	}
	return nil
}
