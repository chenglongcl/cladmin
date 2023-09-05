package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysBulletin = "sys_bulletin"

// SysBulletin mapped from table <sys_bulletin>
type SysBulletin struct {
	ID        uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Title     string         `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Tag       string         `gorm:"column:tag;type:varchar(255);not null" json:"tag"`
	Content   string         `gorm:"column:content;type:text;not null" json:"content"`
	CreatedAt *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"` // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysBulletin's table name
func (*SysBulletin) TableName() string {
	return TableNameSysBulletin
}
