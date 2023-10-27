package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysDept = "sys_dept"

// SysDept mapped from table <sys_dept>
type SysDept struct {
	ID        uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"` // id
	ParentID  uint64         `gorm:"column:parent_id;type:int(11) unsigned;not null" json:"parentId"`         // 上级ID
	Name      string         `gorm:"column:name;type:varchar(255);not null" json:"name"`                      // 部门名称
	Sort      int32          `gorm:"column:sort;type:tinyint(4);not null" json:"sort"`                        // 排序
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`                       // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysDept's table name
func (*SysDept) TableName() string {
	return TableNameSysDept
}
