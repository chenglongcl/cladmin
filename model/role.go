package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Role struct {
	BaseModel
	RoleName     string `json:"role_name"`
	Remark       string `json:"remark"`
	CreateUserId uint64 `json:"create_user_id"`
	Menu         []Menu `json:"menu" gorm:"many2many:sys_role_menu;"`
}

func (r *Role) TableName() string {
	return viper.GetString("db.prefix") + "role"
}

func AddRole(data map[string]interface{}) (id uint64, err error) {
	role := Role{
		RoleName:     data["role_name"].(string),
		Remark:       data["remark"].(string),
		CreateUserId: data["create_user_id"].(uint64),
	}
	var menu []Menu
	DB.Self.Where("id in (?)", data["menu_id_list"].([]int64)).Find(&menu)
	if err := DB.Self.Create(&role).Association("Menu").Append(menu).Error; err != nil {
		return 0, err
	}
	return role.Id, nil
}

func GetRole(id uint64) (*Role, error) {
	var role Role
	err := DB.Self.Preload("Menu").Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &role, nil
}

func GetRolesAll() ([]*Role, error) {
	var role []*Role
	err := DB.Self.Preload("Menu").Find(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return role, nil
}
