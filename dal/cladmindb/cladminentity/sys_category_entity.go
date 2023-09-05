package cladminentity

import "sync"

type CategoryInfo struct {
	Id         uint64 `json:"categoryId"`
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
