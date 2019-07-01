package model

import "github.com/spf13/viper"

type UserRole struct {
	BaseModel
	UserID uint64 `gorm:"column:user_id"`
	RoleID uint64 `gorm:"column:role_id"`
}

func (m *UserRole) TableName() string {
	return viper.GetString("db.prefix") + "user_role"
}

func DeleteUserRoleByRoleID(id uint64) error {
	var userRole UserRole
	return SelectDB("self").Unscoped().Where("role_id = ?", id).Delete(&userRole).Error
}
