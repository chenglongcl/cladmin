package cladminmodel

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysUserToken = "sys_user_token"

// SysUserToken mapped from table <sys_user_token>
type SysUserToken struct {
	UserID      uint64         `gorm:"column:user_id;type:int(11) unsigned;primaryKey" json:"userId"`
	Token       string         `gorm:"column:token;type:varchar(500);not null" json:"token"`          // token
	ExpireTime  time.Time      `gorm:"column:expire_time;type:datetime;not null" json:"expireTime"`   // 过期时间
	RefreshTime time.Time      `gorm:"column:refresh_time;type:datetime;not null" json:"refreshTime"` // 更新时间
	CreatedAt   *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`             // 创建时间
	UpdatedAt   *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysUserToken's table name
func (*SysUserToken) TableName() string {
	return TableNameSysUserToken
}
