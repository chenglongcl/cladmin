package cladminmodel

import (
	"gorm.io/gorm"
	"time"
)

const TableNameSysArticle = "sys_article"

// SysArticle mapped from table <sys_article>
type SysArticle struct {
	ID          uint64         `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	UserID      uint64         `gorm:"column:user_id;type:int(11) unsigned;not null" json:"userId"`
	CateID      uint64         `gorm:"column:cate_id;type:int(11) unsigned;not null" json:"cateId"`
	Title       string         `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Thumb       string         `gorm:"column:thumb;type:varchar(255);not null" json:"thumb"`
	Content     string         `gorm:"column:content;type:text;not null" json:"content"`
	ReleaseTime string         `gorm:"column:release_time;type:varchar(255);not null" json:"releaseTime"`
	CreatedAt   *time.Time     `gorm:"column:created_at;type:timestamp" json:"createdAt"` // 创建时间
	UpdatedAt   *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`
}

// TableName SysArticle's table name
func (*SysArticle) TableName() string {
	return TableNameSysArticle
}
