package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysUserRole = "sys_user_role"

// SysUserRole mapped from table <sys_user_role>
type SysUserRole struct {
	ID        uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	UserID    uint64         `gorm:"column:user_id;type:int(11) unsigned;not null" json:"userId"` // 用户ID
	RoleID    uint64         `gorm:"column:role_id;type:int(11) unsigned;not null" json:"roleId"` // 角色ID
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`           // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysUserRole's table name
func (*SysUserRole) TableName() string {
	return TableNameSysUserRole
}
