package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysMenu = "sys_menu"

// SysMenu mapped from table <sys_menu>
type SysMenu struct {
	ID        uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	ParentID  uint64         `gorm:"column:parent_id;type:int(11) unsigned;not null" json:"parentId"` // 父菜单ID，一级菜单为0
	Name      string         `gorm:"column:name;type:varchar(50);not null" json:"name"`               // 菜单名称
	URL       string         `gorm:"column:url;type:varchar(200);not null" json:"url"`                // 菜单URL
	Perms     string         `gorm:"column:perms;type:varchar(500);not null" json:"perms"`            // 授权(多个用逗号分隔，如：user:list,user:create)
	Type      int64          `gorm:"column:type;type:int(11);not null" json:"type"`                   // 类型   0：目录   1：菜单   2：按钮
	Icon      string         `gorm:"column:icon;type:varchar(50);not null" json:"icon"`               // 菜单图标
	OrderNum  int64          `gorm:"column:order_num;type:int(11);not null" json:"orderNum"`          // 排序
	IsTab     bool           `gorm:"column:is_tab;type:tinyint(1);not null" json:"isTab"`
	Status    bool           `gorm:"column:status;type:tinyint(1);not null" json:"status"`
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"` // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
	Roles     []*SysRole     `gorm:"many2many:sys_role_menu;foreignKey:id;joinForeignKey:menu_id;references:id;joinReferences:role_id;" json:"roles"`
}

// TableName SysMenu's table name
func (*SysMenu) TableName() string {
	return TableNameSysMenu
}
