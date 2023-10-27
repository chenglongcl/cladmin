package cladminentity

import "sync"

type DictDataInfo struct {
	ID         uint64 `json:"dictDataId"`
	DictTypeID uint64 `json:"dictTypeId"`
	DictLabel  string `json:"dictLabel"`
	DictValue  string `json:"dictValue"`
	Remark     string `json:"remark"`
	Sort       uint64 `json:"sort"`
	CreateTime string `json:"createTime"`
}

type DictDataList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*DictDataInfo
}
