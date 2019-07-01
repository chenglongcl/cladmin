package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

type Category struct {
	BaseModel
	ParentID uint64 `gorm:"column:parent_id"`
	Name     string `gorm:"column:name"`
	Icon     string `gorm:"column:icon"`
	OrderNum int64  `gorm:"column:order_num"`
}

type CategoryInfo struct {
	Id         uint64 `json:"categoryId"`
	ParentID   uint64 `json:"parentId"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	OrderNum   int64  `json:"orderNum"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type CategoryList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*CategoryInfo
}

func (c *Category) TableName() string {
	return viper.GetString("db.prefix") + "category"
}

func AddCategory(data map[string]interface{}) error {
	category := Category{
		ParentID: data["parent_id"].(uint64),
		Name:     data["name"].(string),
		Icon:     data["icon"].(string),
		OrderNum: data["order_num"].(int64),
	}
	if err := SelectDB("self").Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func EditCategory(data map[string]interface{}) error {
	var category Category
	if err := SelectDB("self").Model(&category).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func GetCategory(id uint64) (*Category, error) {
	var category Category
	err := SelectDB("self").Where("id = ?", id).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &category, nil
}

func GetCategoryList(w map[string]interface{}) ([]*Category, error) {
	var categoryList []*Category
	where, values, err := WhereBuild(w)
	if err != nil {
		return nil, err
	}
	if err := SelectDB("self").Where(where, values...).Order("parent_id asc,order_num asc").
		Find(&categoryList).Error; err != nil {
		return nil, err
	}
	return categoryList, nil
}

func DeleteCategory(id uint64) error {
	var category Category
	if err := SelectDB("self").Where("id = ?", id).Unscoped().Delete(&category).Error; err != nil {
		return err
	}
	return nil
}
