package cladminmodel

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysRoleMenu = "sys_role_menu"

// SysRoleMenu mapped from table <sys_role_menu>
type SysRoleMenu struct {
	ID        uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	RoleID    uint64         `gorm:"column:role_id;type:int(11) unsigned;not null" json:"roleId"` // 角色ID
	MenuID    uint64         `gorm:"column:menu_id;type:int(11) unsigned;not null" json:"menuId"` // 菜单ID
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`           // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysRoleMenu's table name
func (*SysRoleMenu) TableName() string {
	return TableNameSysRoleMenu
}
