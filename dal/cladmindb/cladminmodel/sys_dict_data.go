package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysDictData = "sys_dict_data"

// SysDictData mapped from table <sys_dict_data>
type SysDictData struct {
	ID         uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"` // id
	DictTypeID uint64         `gorm:"column:dict_type_id;type:int(11) unsigned;not null" json:"dictTypeId"`    // 字典类型ID
	DictLabel  string         `gorm:"column:dict_label;type:varchar(255);not null" json:"dictLabel"`           // 字典标签
	DictValue  string         `gorm:"column:dict_value;type:varchar(255);not null" json:"dictValue"`           // 字典值
	Remark     string         `gorm:"column:remark;type:varchar(255);not null" json:"remark"`                  // 备注
	Sort       uint64         `gorm:"column:sort;type:int(11) unsigned;not null" json:"sort"`                  // 排序
	CreatedAt  *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`                       // 创建时间
	UpdatedAt  *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysDictData's table name
func (*SysDictData) TableName() string {
	return TableNameSysDictData
}
