package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

type Article struct {
	BaseModel
	UserId  uint64 `gorm:"column:user_id"`
	CateId  uint64 `gorm:"column:cate_id"`
	Title   string `gorm:"column:title;"`
	Content string `gorm:"column:content;type:text"`
	Thumb   string `gorm:"column:thumb;type:text"`
}

type ArticleInfo struct {
	Id         uint64   `json:"articleId"`
	UserId     uint64   `json:"userId"`
	CateId     uint64   `json:"cateId"`
	Title      string   `json:"title"`
	Thumb      []string `json:"thumb"`
	CreateTime string   `json:"createTime"`
}

type ArticleList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*ArticleInfo
}

func (a *Article) TableName() string {
	return viper.GetString("db.prefix") + "article"
}

func AddArticle(data map[string]interface{}) error {
	article := Article{
		UserId:  data["user_id"].(uint64),
		CateId:  data["cate_id"].(uint64),
		Title:   data["title"].(string),
		Content: data["content"].(string),
		Thumb:   data["thumb"].(string),
	}
	if err := DB.Self.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func EditArticle(data map[string]interface{}) error {
	var article Article
	if err := DB.Self.Model(&article).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func GetArticle(id uint64) (*Article, error) {
	var article Article
	err := DB.Self.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func GetArticleList(w map[string]interface{}, offset, limit uint64) ([]*Article, uint64, error) {
	articles := make([]*Article, 0)
	var count uint64
	where, values, _ := WhereBuild(w)
	if err := DB.Self.Model(&Article{}).Where(where, values...).Count(&count).Error; err != nil {
		return articles, count, err
	}
	if err := DB.Self.Model(&Article{}).Where(where, values...).Offset(offset).Limit(limit).Order("id desc").
		Find(&articles).Error;
		err != nil {
		return articles, count, err
	}
	return articles, count, nil
}

func DeleteArticle(id uint64) error {
	var article Article
	if err := DB.Self.Where("id = ?", id).Delete(&article).Error; err != nil {
		return err
	}
	return nil
}
