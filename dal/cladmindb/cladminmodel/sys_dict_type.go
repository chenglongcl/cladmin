package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysDictType = "sys_dict_type"

// SysDictType mapped from table <sys_dict_type>
type SysDictType struct {
	ID        uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"` // id
	DictType  string         `gorm:"column:dict_type;type:varchar(100);not null" json:"dictType"`             // 字典类型
	DictName  string         `gorm:"column:dict_name;type:varchar(255);not null" json:"dictName"`             // 字典名称
	Remark    string         `gorm:"column:remark;type:varchar(255);not null" json:"remark"`                  // 备注
	Sort      uint64         `gorm:"column:sort;type:int(11) unsigned;not null" json:"sort"`                  // 排序
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`                       // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysDictType's table name
func (*SysDictType) TableName() string {
	return TableNameSysDictType
}
