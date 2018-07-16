package model

import (
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type ArticleModel struct {
	BaseModel
	CateId  uint64      `json:"cate_id" gorm:"column:cate_id;not null"`
	Title   string      `json:"title" gorm:"column:title;not null" binding:"required" validate:"min=1,max=32"`
	Content string      `json:"content" gorm:"column:content;type:text"`
	Images  string      `json:"images" gorm:"column:images;type:text"`
	Uid     uint64      `json:"uid" gorm:"column:uid;not null"`
	Author  AuthorModel `gorm:"ForeignKey:Uid"`
}

func (a *ArticleModel) TableName() string {
	return viper.GetString("db.prefix") + "articles"
}
func (a *ArticleModel) CreateArticle() error {
	return DB.Self.Create(&a).Error
}
func GetArticle(id uint64) (*ArticleModel, error) {
	a := &ArticleModel{}
	d := DB.Self.Where("id = ?", id).First(&a)
	DB.Self.Model(&a).Select("id,username").Related(&a.Author, "Author")
	return a, d.Error
}
func ListArticle(CateId uint64, Offset uint64, Limit uint64) ([]*ArticleModel, uint64, error) {
	articles := make([]*ArticleModel, 0)
	var count uint64
	where := map[string]interface{}{"cate_id": CateId}
	if err := DB.Self.Model(&ArticleModel{}).Where(where).Count(&count).Error; err != nil {
		return articles, count, err
	}
	if err := DB.Self.Preload("Author").Where(where).Offset(Offset).Limit(Limit).
		Order("id desc").Find(&articles).Error; err != nil {
		return articles, count, err
	}
	return articles, count, nil
}
func (a *ArticleModel) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}