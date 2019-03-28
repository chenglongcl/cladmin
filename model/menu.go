package model

import (
	"github.com/spf13/viper"
)

type Menu struct {
	BaseModel
	ParentId uint64 `gorm:"column:parent_id"`
	Name     string `gorm:"column:name"`
	Url      string `gorm:"column:url"`
	Perms    string `gorm:"column:perms"`
	Type     int64  `gorm:"column:type"`
	Icon     string `gorm:"column:icon"`
	OrderNum int64  `gorm:"column:order_num"`
}

func (m *Menu) TableName() string {
	return viper.GetString("db.prefix") + "menu"
}

func AddMenu(data map[string]interface{}) error {
	menu := Menu{
		ParentId: data["parent_id"].(uint64),
		Name:     data["name"].(string),
		Url:      data["url"].(string),
		Perms:    data["perms"].(string),
		Type:     data["type"].(int64),
		Icon:     data["icon"].(string),
		OrderNum: data["order_num"].(int64),
	}
	if err := DB.Self.Create(&menu).Error; err != nil {
		return err
	}
	return nil
}

func EditMenu(data map[string]interface{}) error {
	var menu Menu
	if err := DB.Self.Model(&menu).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditMenuGetRoles(id uint64) []uint64 {
	var (
		menu Menu
		role []Role
	)
	DB.Self.Model(&menu).Where("c.id = ?", id).
		Joins(" left join sys_role_menu b on sys_role.id=b.role_id left join sys_menu c on c.id=b.menu_id").
		Find(&role)
	var roleList []uint64
	for _, v := range role {
		roleList = append(roleList, v.Id)
	}
	return roleList
}
