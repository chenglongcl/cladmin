package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}
type Token struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}
type ArticleInfo struct {
	Id        uint64   `json:"id"`
	Uid       uint64   `json:"uid"`
	CateId    uint64   `json:"cate_id"`
	Title     string   `json:"title"`
	Images    []string `json:"images"`
	Author    Author
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type Author struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}
type ArticleList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*ArticleInfo
}
