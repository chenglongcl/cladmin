package cladminentity

import "sync"

type DictTypeInfo struct {
	ID         uint64          `json:"dictTypeId"`
	DictType   string          `json:"dictType"`   // 字典类型
	DictName   string          `json:"dictName"`   // 字典名称
	Remark     string          `json:"remark"`     // 备注
	Sort       uint64          `json:"sort"`       // 排序
	CreateTime string          `json:"createTime"` // 创建时间
	DataList   []*DictDataInfo `json:"dataList,omitempty"`
}

type DictTypeList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*DictTypeInfo
}
