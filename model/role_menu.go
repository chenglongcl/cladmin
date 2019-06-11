package model

import "github.com/spf13/viper"

type RoleMenu struct {
	BaseModel
	RoleId uint64
	MenuId uint64
}

func (m *RoleMenu) TableName() string {
	return viper.GetString("db.prefix") + "role_menu"
}

func DeleteRoleMenuByMenuId(id uint64) error {
	var roleMenu RoleMenu
	return SelectDB("self").Unscoped().Where("menu_id = ?", id).Delete(&roleMenu).Error
}
