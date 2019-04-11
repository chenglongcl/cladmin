package model

import "github.com/spf13/viper"

type UserRole struct {
	BaseModel
	UserId uint64
	RoleId uint64
}

func (m *UserRole) TableName() string {
	return viper.GetString("db.prefix") + "user_role"
}

func DeleteUserRoleByRoleId(id uint64) error {
	var userRole UserRole
	return DB.Self.Unscoped().Where("role_id = ?", id).Delete(&userRole).Error
}
