package menuservice

import (
	"cladmin/model"
	"cladmin/pkg/errno"
	"github.com/casbin/casbin"
	"strings"
	"sync"
)

type Menu struct {
	ID       uint64
	ParentID uint64
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
		"parent_id": a.ParentID,
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
	menu, err := model.GetMenu(a.ID)
	if err != nil {
		return menu, errno.ErrDatabase
	}
	return menu, nil
}

func (a *Menu) GetList(w map[string]interface{}) ([]*model.MenuInfo, *errno.Errno) {
	menus, err := model.GetMenuList(w)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	var ids []uint64
	for _, menu := range menus {
		ids = append(ids, menu.ID)
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
			menuList.IdMap[menu.ID] = &model.MenuInfo{
				ID:         menu.ID,
				ParentID:   menu.ParentID,
				ParentName: "",
				Name:       menu.Name,
				Url:        menu.Url,
				Perms:      menu.Perms,
				Type:       menu.Type,
				Icon:       menu.Icon,
				Open:       0,
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
	if menuExists, _ := model.CheckMenuByID(a.ID); !menuExists {
		return nil, errno.ErrRecordNotFound
	}
	data := map[string]interface{}{
		"id":        a.ID,
		"parent_id": a.ParentID,
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
	roleList := model.EditMenuGetRoles(a.ID)
	return roleList, nil
}

func (a *Menu) Delete() ([]uint64, *errno.Errno) {
	roleList := model.EditMenuGetRoles(a.ID)
	children, err := model.GetMenuList(map[string]interface{}{
		"parent_id": a.ID,
	})
	if err != nil {
		return nil, errno.ErrDatabase
	}
	if len(children) > 0 {
		return nil, errno.ErrRecordHasChildren
	}
	if err := model.DeleteMenu(a.ID); err != nil {
		return nil, errno.ErrDatabase
	}
	return roleList, nil
}

func (a *Menu) GetMenuNavByUserID(userID uint64) ([]*MenuTree, []string, *errno.Errno) {
	var (
		menus   []*model.Menu
		modeErr error
	)
	if userID == 1 {
		//admin
		w := map[string]interface{}{}
		menus, modeErr = model.GetMenuList(w)
	} else {
		menus, modeErr = model.GetMenuNavByUserId(userID)
	}
	if modeErr != nil {
		return nil, nil, errno.ErrDatabase
	}
	var (
		menuTrees MenuTrees
	)
	permissions := make([]string, 0)
	for _, menu := range menus {
		if menu.Type != 2 {
			menuTrees = append(menuTrees, &MenuTree{
				MenuID:     menu.ID,
				Name:       menu.Name,
				Open:       0,
				OrderNum:   menu.OrderNum,
				Icon:       menu.Icon,
				Url:        menu.Url,
				ParentID:   menu.ParentID,
				ParentName: "",
				Perms:      menu.Perms,
				Type:       menu.Type,
			})
		}
		if menu.Perms != "" {
			pSlice := strings.Split(menu.Perms, ",")
			permissions = append(permissions, pSlice...)
		}
	}
	list := menuTrees.ToTree()
	return list, permissions, nil
}

// MenuTree 菜单树
type MenuTree struct {
	MenuID     uint64       `json:"menuId"`
	Name       string       `json:"name"`
	Open       int64        `json:"open"`
	OrderNum   int64        `json:"orderNum"`
	Icon       string       `json:"icon"`
	Url        string       `json:"url"`
	ParentID   uint64       `json:"parentId"`
	ParentName string       `json:"parentName"`
	Perms      string       `json:"perms"`
	Type       int64        `json:"type"`
	List       *[]*MenuTree `json:"list,omitempty"`
}

// MenuTrees 菜单树列表
type MenuTrees []*MenuTree

// ForEach 遍历数据项
func (a MenuTrees) ForEach(fn func(*MenuTree, int)) MenuTrees {
	for i, item := range a {
		fn(item, i)
	}
	return a
}

// ToTree 转换为树形结构
func (a MenuTrees) ToTree() []*MenuTree {
	mi := make(map[uint64]*MenuTree)
	for _, item := range a {
		mi[item.MenuID] = item
	}

	var list []*MenuTree
	for _, item := range a {
		if item.ParentID == 0 {
			list = append(list, item)
			continue
		}
		if pitem, ok := mi[item.ParentID]; ok {
			if pitem.List == nil {
				var children []*MenuTree
				children = append(children, item)
				pitem.List = &children
				continue
			}
			*pitem.List = append(*pitem.List, item)
		}
	}
	return list
}
