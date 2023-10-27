package cladminentity

import "sync"

type ConfigInfo struct {
	ID         uint64 `json:"configId"`
	ParamKey   string `json:"paramKey"`
	ParamValue string `json:"paramValue"`
	Type       int32  `json:"type"`
	Status     bool   `json:"status"`
	Remark     string `json:"remark"`
	Locked     bool   `json:"locked"`
}

type ConfigList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*ConfigInfo
}
