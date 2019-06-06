package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

type PublicNotice struct {
	BaseModel
	Title   string `gorm:"column:title"`
	Tag     string `gorm:"column:tag"`
	Content string `gorm:"column:content;type:text"`
}

type PublicNoticeInfo struct {
	Id         uint64 `json:"id"`
	Title      string `json:"title"`
	Tag        string `json:"tag"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

type PublicNoticeList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*PublicNoticeInfo
}

func (a *PublicNotice) TableName() string {
	return viper.GetString("db.prefix") + "public_notice"
}

func AddPublicNotice(data map[string]interface{}) error {
	publicNotice := &PublicNotice{
		Title:   data["title"].(string),
		Tag:     data["tag"].(string),
		Content: data["content"].(string),
	}
	if err := DB.Self.Create(&publicNotice).Error; err != nil {
		return err
	}
	return nil
}

func EditPublicNotice(data map[string]interface{}) error {
	var publicNotice PublicNotice
	if err := DB.Self.Model(&publicNotice).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func GetPublicNotice(id uint64) (*PublicNotice, error) {
	var publicNotice PublicNotice
	err := DB.Self.Where("id = ?", id).First(&publicNotice).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &publicNotice, nil
}

func GetPublicNoticeList(w map[string]interface{}, offset, limit uint64) ([]*PublicNotice, uint64, error) {
	publicNotices := make([]*PublicNotice, 0)
	var count uint64
	where, values, _ := WhereBuild(w)
	if err := DB.Self.Model(&PublicNotice{}).Where(where, values...).Count(&count).Error; err != nil {
		return publicNotices, count, err
	}
	if err := DB.Self.Where(where, values...).Offset(offset).Limit(limit).Order("id desc").
		Find(&publicNotices).Error; err != nil {
		return publicNotices, count, err
	}
	return publicNotices, count, nil
}

func DeletePublicNotice(id uint64) error {
	var publicNotice PublicNotice
	if err := DB.Self.Where("id = ?", id).Delete(&publicNotice).Error; err != nil {
		return err
	}
	return nil
}
