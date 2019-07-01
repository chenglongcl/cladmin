package model

import "github.com/spf13/viper"

type RoleMenu struct {
	BaseModel
	RoleID uint64 `gorm:"column:role_id"`
	MenuID uint64 `gorm:"column:menu_id"`
}

func (m *RoleMenu) TableName() string {
	return viper.GetString("db.prefix") + "role_menu"
}

func DeleteRoleMenuByMenuID(id uint64) error {
	var roleMenu RoleMenu
	return SelectDB("self").Unscoped().Where("menu_id = ?", id).Delete(&roleMenu).Error
}
