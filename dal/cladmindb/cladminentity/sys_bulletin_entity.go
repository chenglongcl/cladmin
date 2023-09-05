package cladminentity

import "sync"

type BulletinInfo struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	Tag        string `json:"tag"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

type BulletinList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*BulletinInfo
}
