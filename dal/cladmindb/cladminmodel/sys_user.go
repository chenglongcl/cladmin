package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	ID           uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Username     string         `gorm:"column:username;type:varchar(50);not null" json:"username"`                // 用户名
	Password     string         `gorm:"column:password;type:varchar(100);not null" json:"password"`               // 密码
	Email        string         `gorm:"column:email;type:varchar(100);not null" json:"email"`                     // 邮箱
	Mobile       string         `gorm:"column:mobile;type:varchar(100);not null" json:"mobile"`                   // 手机号
	Gender       int32          `gorm:"column:gender;type:tinyint(4);not null" json:"gender"`                     // 性别
	DeptID       uint64         `gorm:"column:dept_id;type:int(11) unsigned;not null" json:"deptId"`              // 部门ID
	Status       int32          `gorm:"column:status;type:tinyint(4);not null" json:"status"`                     // 状态  0：禁用   1：正常
	SuperAdmin   bool           `gorm:"column:super_admin;type:tinyint(1);not null" json:"superAdmin"`            // 超级管理员   0：否   1：是
	CreateUserID uint64         `gorm:"column:create_user_id;type:int(11) unsigned;not null" json:"createUserId"` // 创建者ID
	CreatedAt    *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`                        // 创建时间
	UpdatedAt    *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
	Roles        []*SysRole     `gorm:"many2many:sys_user_role;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:role_id;" json:"roles"`
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
