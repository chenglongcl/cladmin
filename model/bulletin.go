package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

type Bulletin struct {
	BaseModel
	Title   string `gorm:"column:title"`
	Tag     string `gorm:"column:tag"`
	Content string `gorm:"column:content;type:text"`
}

type BulletinInfo struct {
	Id         uint64 `json:"id"`
	Title      string `json:"title"`
	Tag        string `json:"tag"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

type BulletinList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*BulletinInfo
}

func (a *Bulletin) TableName() string {
	return viper.GetString("db.prefix") + "bulletin"
}

func AddBulletin(data map[string]interface{}) error {
	bulletin := &Bulletin{
		Title:   data["title"].(string),
		Tag:     data["tag"].(string),
		Content: data["content"].(string),
	}
	if err := SelectDB("self").Create(&bulletin).Error; err != nil {
		return err
	}
	return nil
}

func EditBulletin(data map[string]interface{}) error {
	var bulletin Bulletin
	if err := SelectDB("self").Model(&bulletin).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func GetBulletin(id uint64) (*Bulletin, error) {
	var bulletin Bulletin
	err := SelectDB("self").Where("id = ?", id).First(&bulletin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &bulletin, nil
}

func GetBulletinList(w map[string]interface{}, offset, limit uint64) ([]*Bulletin, uint64, error) {
	bulletin := make([]*Bulletin, 0)
	var count uint64
	where, values, _ := WhereBuild(w)
	if err := SelectDB("self").Model(&Bulletin{}).Where(where, values...).Count(&count).Error; err != nil {
		return bulletin, count, err
	}
	if err := SelectDB("self").Where(where, values...).Offset(offset).Limit(limit).Order("id desc").
		Find(&bulletin).Error; err != nil {
		return bulletin, count, err
	}
	return bulletin, count, nil
}

func DeleteBulletin(id uint64) error {
	var bulletin Bulletin
	if err := SelectDB("self").Where("id = ?", id).Delete(&bulletin).Error; err != nil {
		return err
	}
	return nil
}
