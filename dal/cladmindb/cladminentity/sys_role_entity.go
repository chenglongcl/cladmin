package cladminentity

import "sync"

type RoleInfo struct {
	ID           uint64  `json:"roleId"`
	RoleName     string  `json:"roleName"`
	Remark       string  `json:"remark"`
	MenuIDList   []int64 `json:"menuIdList"`
	CreateUserID uint64  `json:"createUserId"`
	CreateTime   string  `json:"createTime"`
}

type RoleList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*RoleInfo
}
