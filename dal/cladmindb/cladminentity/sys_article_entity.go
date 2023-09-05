package cladminentity

import "sync"

type ArticleInfo struct {
	ID          uint64 `json:"articleId"`
	UserID      uint64 `json:"userId"`
	CateID      uint64 `json:"cateId"`
	Title       string `json:"title"`
	Thumb       string `json:"thumb"`
	ReleaseTime string `json:"releaseTime"`
}

type ArticleList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*ArticleInfo
}
