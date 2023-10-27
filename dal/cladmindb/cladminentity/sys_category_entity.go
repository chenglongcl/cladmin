package cladminentity

import "sync"

type CategoryInfo struct {
	ID         uint64 `json:"categoryId"`
	ParentID   uint64 `json:"parentId"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	OrderNum   int64  `json:"orderNum"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type CategoryList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*CategoryInfo
}

type CategoryTree struct {
	ID         uint64          `json:"categoryId"`
	ParentID   uint64          `json:"parentId"`
	Name       string          `json:"name"`
	Icon       string          `json:"icon"`
	OrderNum   int64           `json:"orderNum"`
	CreateTime string          `json:"createTime"`
	UpdateTime string          `json:"updateTime"`
	Children   []*CategoryTree `json:"children,omitempty"`
}

func (d *CategoryTree) IsEqual(father *CategoryTree, child *CategoryTree) bool {
	return father.ID == child.ParentID
}
func (d *CategoryTree) SetChildren(father *CategoryTree, children []*CategoryTree) {
	father.Children = children
}
func (d *CategoryTree) RetFather(father *CategoryTree) bool {
	return father.ParentID == 0
}
