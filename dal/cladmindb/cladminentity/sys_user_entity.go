package cladminentity

import "sync"

type UserInfo struct {
	ID           uint64 `json:"userId"`
	Username     string `json:"username"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
	Gender       int32  `json:"gender"`
	SuperAdmin   bool   `json:"superAdmin"`
	Status       int32  `json:"status"`
	CreateUserID uint64 `json:"createUserId"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}
