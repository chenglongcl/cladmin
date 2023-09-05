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
	Open       int64  `json:"open"`
	OrderNum   int64  `json:"orderNum"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type MenuList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*MenuInfo
}
