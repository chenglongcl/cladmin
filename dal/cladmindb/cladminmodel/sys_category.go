package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysCategory = "sys_category"

// SysCategory mapped from table <sys_category>
type SysCategory struct {
	ID        uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	ParentID  uint64         `gorm:"column:parent_id;type:int(11) unsigned;not null" json:"parentId"` // 父菜单ID，一级菜单为0
	Name      string         `gorm:"column:name;type:varchar(50);not null" json:"name"`               // 菜单名称
	Icon      string         `gorm:"column:icon;type:varchar(50);not null" json:"icon"`               // 菜单图标
	OrderNum  int64          `gorm:"column:order_num;type:int(11);not null" json:"orderNum"`          // 排序
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`               // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysCategory's table name
func (*SysCategory) TableName() string {
	return TableNameSysCategory
}
