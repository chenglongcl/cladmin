package model

import "github.com/spf13/viper"

type Menu struct {
	BaseModel
	ParentId int64  `json:"parent_id" gorm:"column:parent_id"`
	Name     string `json:"name" gorm:"column:name"`
	Url      string `json:"url" gorm:"column:url"`
	Perms    string `json:"perms" gorm:"column:perms"`
	Type     int64  `json:"type" gorm:"column:type"`
	Icon     string `json:"icon" gorm:"column:icon"`
	OrderNum int64  `json:"order_num" gorm:"column:order_num"`
}

func (m *Menu) TableName() string {
	return viper.GetString("db.prefix") + "menu"
}

func (m *Menu) CreateMenu() error {
	return DB.Self.Create(&m).Error
}
