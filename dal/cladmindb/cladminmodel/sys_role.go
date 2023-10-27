package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysRole = "sys_role"

// SysRole mapped from table <sys_role>
type SysRole struct {
	ID           uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	RoleName     string         `gorm:"column:role_name;type:varchar(100);not null" json:"roleName"`              // 角色名称
	Remark       string         `gorm:"column:remark;type:varchar(100);not null" json:"remark"`                   // 备注
	MenuIDList   string         `gorm:"column:menu_id_list;type:text;not null" json:"menuIdList"`                 // 配合前端tree半选 666666为临时KEY分隔符
	CreateUserID uint64         `gorm:"column:create_user_id;type:int(11) unsigned;not null" json:"createUserId"` // 创建者ID
	CreatedAt    *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`                        // 创建时间
	UpdatedAt    *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
	Menus        []*SysMenu     `gorm:"many2many:sys_role_menu;foreignKey:id;joinForeignKey:role_id;references:id;joinReferences:menu_id;" json:"menus"`
	Users        []*SysUser     `gorm:"many2many:sys_user_role;foreignKey:id;joinForeignKey:role_id;references:id;joinReferences:user_id;" json:"users"`
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}
