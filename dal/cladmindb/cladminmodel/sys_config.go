package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysConfig = "sys_config"

// SysConfig mapped from table <sys_config>
type SysConfig struct {
	ID         uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	ParamKey   string         `gorm:"column:param_key;type:varchar(255);not null" json:"paramKey"` // key
	ParamValue string         `gorm:"column:param_value;type:longtext;not null" json:"paramValue"` // value
	Type       int32          `gorm:"column:type;type:tinyint(4);not null;default:1" json:"type"`  // 1字符串值类型 2字符串JSON类型
	Status     bool           `gorm:"column:status;type:tinyint(1);not null" json:"status"`        // 状态   0：隐藏   1：显示
	Remark     string         `gorm:"column:remark;type:varchar(255);not null" json:"remark"`      // 备注
	Locked     bool           `gorm:"column:locked;type:tinyint(1);not null" json:"locked"`        // 锁定  0：否  1：是
	CreatedAt  *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"`           // 创建时间
	UpdatedAt  *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysConfig's table name
func (*SysConfig) TableName() string {
	return TableNameSysConfig
}
