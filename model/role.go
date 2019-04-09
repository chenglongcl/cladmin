package model

import (
	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"github.com/spf13/viper"
	"sync"
)

type Role struct {
	BaseModel
	RoleName     string
	Remark       string
	MenuIdList   string
	CreateUserId uint64
	Menu         []Menu `gorm:"many2many:sys_role_menu;"`
}

type RoleInfo struct {
	Id           uint64  `json:"roleId"`
	RoleName     string  `json:"roleName"`
	Remark       string  `json:"remark"`
	MenuIdList   []int64 `json:"menuIdList"`
	CreateUserId uint64  `json:"createUserId"`
	CreateTime   string  `json:"createTime"`
}

type RoleList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*RoleInfo
}

func (r *Role) TableName() string {
	return viper.GetString("db.prefix") + "role"
}

func CheckRoleById(id uint64) (bool, error) {
	var role Role
	err := DB.Self.Select("id").Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.Id > 0 {
		return true, nil
	}
	return false, nil
}

func CheckRoleByRoleName(roleName string) (bool, error) {
	var role Role
	err := DB.Self.Select("id").Where("role_name = ?", roleName).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.Id > 0 {
		return true, nil
	}
	return false, nil
}

func CheckRoleByRoleNameId(id uint64, roleName string) (bool, error) {
	var role Role
	err := DB.Self.Select("id").Where("role_name = ? AND id != ?", roleName, id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.Id > 0 {
		return true, nil
	}
	return false, nil
}

func AddRole(data map[string]interface{}) (id uint64, err error) {
	menuIdListJson, _ := jsoniter.MarshalToString(data["menu_id_list"])
	role := Role{
		RoleName:     data["role_name"].(string),
		Remark:       data["remark"].(string),
		CreateUserId: data["create_user_id"].(uint64),
		MenuIdList:   menuIdListJson,
	}
	var menu []Menu
	DB.Self.Where("id in (?)", data["menu_id_list"].([]int64)).Find(&menu)
	if err := DB.Self.Create(&role).Association("Menu").Append(menu).Error; err != nil {
		return 0, err
	}
	return role.Id, nil
}

func EditRole(data map[string]interface{}) error {
	var (
		role Role
		menu []Menu
	)
	if err := DB.Self.Where("id = ?", data["id"].(uint64)).First(&role).Error; err != nil {
		return err
	}
	DB.Self.Where("id in (?)", data["menu_id_list"].([]int64)).Find(&menu)
	//delete(data, "menu_id_list")
	data["menu_id_list"], _ = jsoniter.MarshalToString(data["menu_id_list"])
	DB.Self.Model(&role).Association("Menu").Replace(&menu)
	DB.Self.Model(&role).Update(data)
	return nil
}

func DeleteRole(id uint64) error {
	var role Role
	DB.Self.Where("id = ?", id).Preload("Menu").First(&role)
	DB.Self.Model(&role).Association("Menu").Delete(role.Menu)
	if err := DB.Self.Unscoped().Where("id = ?", id).Delete(&role).Error; err != nil {
		return err
	}
	return nil
}

func GetRoleList(w map[string]interface{}, offset, limit uint64) ([]*Role, uint64, error) {
	roles := make([]*Role, 0)
	var count uint64
	where, values, _ := WhereBuild(w)
	if err := DB.Self.Model(&Role{}).Where(where, values...).Count(&count).Error; err != nil {
		return roles, count, err
	}
	if err := DB.Self.Model(&Role{}).Where(where, values...).Offset(offset).Limit(limit).Order("id asc").
		Find(&roles).Error;
		err != nil {
		return roles, count, err
	}
	return roles, count, nil
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
