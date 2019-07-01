package model

import (
	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"github.com/spf13/viper"
	"sync"
)

type Role struct {
	BaseModel
	RoleName     string `gorm:"column:role_name"`
	Remark       string `gorm:"column:remark"`
	MenuIDList   string `gorm:"column:menu_id_list"`
	CreateUserID uint64 `gorm:"column:create_user_id"`
	Menu         []Menu `gorm:"many2many:sys_role_menu;"`
}

type RoleInfo struct {
	Id           uint64  `json:"roleId"`
	RoleName     string  `json:"roleName"`
	Remark       string  `json:"remark"`
	MenuIDList   []int64 `json:"menuIdList"`
	CreateUserID uint64  `json:"createUserId"`
	CreateTime   string  `json:"createTime"`
}

type RoleList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*RoleInfo
}

func (r *Role) TableName() string {
	return viper.GetString("db.prefix") + "role"
}

func CheckRoleByID(id uint64) (bool, error) {
	var role Role
	err := SelectDB("self").Select("id").Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.ID > 0 {
		return true, nil
	}
	return false, nil
}

func CheckRoleByRoleName(roleName string) (bool, error) {
	var role Role
	err := SelectDB("self").Select("id").Where("role_name = ?", roleName).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.ID > 0 {
		return true, nil
	}
	return false, nil
}

func CheckRoleByRoleNameID(id uint64, roleName string) (bool, error) {
	var role Role
	err := SelectDB("self").Select("id").Where("role_name = ? AND id != ?", roleName, id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddRole(data map[string]interface{}) (id uint64, err error) {
	menuIdListJson, _ := jsoniter.MarshalToString(data["menu_id_list"])
	role := Role{
		RoleName:     data["role_name"].(string),
		Remark:       data["remark"].(string),
		CreateUserID: data["create_user_id"].(uint64),
		MenuIDList:   menuIdListJson,
	}
	var menu []Menu
	SelectDB("self").Where("id in (?)", data["menu_id_list"].([]int64)).Find(&menu)
	if err := SelectDB("self").Create(&role).Association("Menu").Append(menu).Error; err != nil {
		return 0, err
	}
	return role.ID, nil
}

func EditRole(data map[string]interface{}) error {
	var (
		role Role
		menu []Menu
	)
	if err := SelectDB("self").Where("id = ?", data["id"].(uint64)).First(&role).Error; err != nil {
		return err
	}
	SelectDB("self").Where("id in (?)", data["menu_id_list"].([]int64)).Find(&menu)
	//delete(data, "menu_id_list")
	//配合前端tree半选状态使用
	data["menu_id_list"], _ = jsoniter.MarshalToString(data["menu_id_list"])
	SelectDB("self").Model(&role).Association("Menu").Replace(&menu)
	SelectDB("self").Model(&role).Update(data)
	return nil
}

func DeleteRole(id uint64) error {
	var role Role
	SelectDB("self").Where("id = ?", id).Preload("Menu").First(&role)
	SelectDB("self").Model(&role).Association("Menu").Delete(role.Menu)
	if err := SelectDB("self").Unscoped().Where("id = ?", id).Delete(&role).Error; err != nil {
		return err
	}
	go func() {
		//删除中间表user_role关联数据
		DeleteUserRoleByRoleID(id)
	}()
	return nil
}

func GetRoleList(w map[string]interface{}, offset, limit uint64) ([]*Role, uint64, error) {
	roles := make([]*Role, 0)
	var count uint64
	where, values, _ := WhereBuild(w)
	if err := SelectDB("self").Model(&Role{}).Where(where, values...).Count(&count).Error; err != nil {
		return roles, count, err
	}
	if err := SelectDB("self").Model(&Role{}).Where(where, values...).Offset(offset).Limit(limit).Order("id asc").
		Find(&roles).Error;
		err != nil {
		return roles, count, err
	}
	return roles, count, nil
}

func GetRole(id uint64) (*Role, error) {
	var role Role
	err := SelectDB("self").Preload("Menu").Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &role, nil
}

func GetRolesAll() ([]*Role, error) {
	var role []*Role
	err := SelectDB("self").Preload("Menu").Find(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return role, nil
}
