package cladminentity

import "sync"

type MenuInfo struct {
	ID         uint64 `json:"menuId"`
	ParentID   uint64 `json:"parentId"`
	ParentName string `json:"parentName"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Perms      string `json:"perms"`
	Type       int64  `json:"type"`
	Icon       string `json:"icon"`
	OrderNum   int64  `json:"orderNum"`
	IsTab      bool   `json:"isTab"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type MenuList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*MenuInfo
}

type MenuTree struct {
	ID         uint64      `json:"menuId"`
	Name       string      `json:"name"`
	OrderNum   int64       `json:"orderNum"`
	IsTab      bool        `json:"isTab"`
	Icon       string      `json:"icon"`
	Url        string      `json:"url"`
	ParentID   uint64      `json:"parentId"`
	ParentName string      `json:"parentName"`
	Perms      string      `json:"perms"`
	Type       int64       `json:"type"`
	Children   []*MenuTree `json:"children,omitempty"`
}

func (d *MenuTree) IsEqual(father *MenuTree, child *MenuTree) bool {
	return father.ID == child.ParentID
}
func (d *MenuTree) SetChildren(father *MenuTree, children []*MenuTree) {
	father.Children = children
}
func (d *MenuTree) RetFather(father *MenuTree) bool {
	return father.ParentID == 0
}
