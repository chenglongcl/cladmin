package menu_service

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"github.com/casbin/casbin"
	"sync"
)

type Menu struct {
	Id       uint64
	ParentId uint64
	Name     string
	Url      string
	Perms    string
	Type     int64
	Icon     string
	OrderNum int64

	Enforcer *casbin.Enforcer `inject:""`
}

func (a *Menu) Add() *errno.Errno {
	data := map[string]interface{}{
		"parent_id": a.ParentId,
		"name":      a.Name,
		"url":       a.Url,
		"perms":     a.Perms,
		"type":      a.Type,
		"icon":      a.Icon,
		"order_num": a.OrderNum,
	}
	if err := model.AddMenu(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func (a *Menu) Get() (*model.Menu, *errno.Errno) {
	menu, err := model.GetMenu(a.Id)
	if err != nil {
		return menu, errno.ErrDatabase
	}
	return menu, nil
}

func (a *Menu) GetList() ([]*model.MenuInfo, *errno.Errno) {
	menus, err := model.GetMenuList()
	if err != nil {
		return nil, errno.ErrDatabase
	}
	var ids []uint64
	for _, menu := range menus {
		ids = append(ids, menu.Id)
	}

	info := make([]*model.MenuInfo, 0)
	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	menuList := model.MenuList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.MenuInfo, len(menus)),
	}
	for _, menu := range menus {
		wg.Add(1)
		go func(menu *model.Menu) {
			defer wg.Done()
			menuList.Lock.Lock()
			defer menuList.Lock.Unlock()
			menuList.IdMap[menu.Id] = &model.MenuInfo{
				Id:         menu.Id,
				ParentId:   menu.ParentId,
				Name:       menu.Name,
				Url:        menu.Url,
				Perms:      menu.Perms,
				Type:       menu.Type,
				Icon:       menu.Icon,
				OrderNum:   menu.OrderNum,
				CreateTime: menu.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdateTime: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(menu)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	}

	for _, id := range ids {
		info = append(info, menuList.IdMap[id])
	}
	return info, nil
}

func (a *Menu) Edit() ([]uint64, *errno.Errno) {
	if menuExists, _ := model.CheckMenuById(a.Id); !menuExists {
		return nil, errno.ErrRecordNotFound
	}
	data := map[string]interface{}{
		"id":        a.Id,
		"parent_id": a.ParentId,
		"name":      a.Name,
		"url":       a.Url,
		"perms":     a.Perms,
		"type":      a.Type,
		"icon":      a.Icon,
		"order_num": a.OrderNum,
	}
	if err := model.EditMenu(data); err != nil {
		return nil, errno.ErrDatabase
	}
	roleList := model.EditMenuGetRoles(a.Id)
	return roleList, nil
}

func (a *Menu) Delete() ([]uint64, *errno.Errno) {
	if err := model.DeleteMenu(a.Id); err != nil {
		return nil, errno.ErrDatabase
	}
	roleList := model.EditMenuGetRoles(a.Id)
	return roleList, nil
}
