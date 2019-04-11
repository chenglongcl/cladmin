package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
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

type MenuInfo struct {
	Id         uint64 `json:"menuId"`
	ParentId   uint64 `json:"parentId"`
	ParentName string `json:"parentName"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Perms      string `json:"perms"`
	Type       int64  `json:"type"`
	Icon       string `json:"icon"`
	Open       int64  `json:"open"`
	OrderNum   int64  `json:"orderNum"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type MenuList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*MenuInfo
}

func (m *Menu) TableName() string {
	return viper.GetString("db.prefix") + "menu"
}

func CheckMenuById(id uint64) (bool, error) {
	var menu Menu
	err := DB.Self.Select("id").Where("id = ?", id).First(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if menu.Id > 0 {
		return true, nil
	}
	return false, nil
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

func GetMenu(id uint64) (*Menu, error) {
	var menu Menu
	err := DB.Self.Where("id = ?", id).First(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &menu, nil
}

func GetMenuList(w map[string]interface{}) ([]*Menu, error) {
	var menuList []*Menu
	where, values, err := WhereBuild(w)
	if err != nil {
		return nil, err
	}
	if err := DB.Self.Where(where, values...).Order("parent_id asc,order_num asc").
		Find(&menuList).Error; err != nil {
		return nil, err
	}
	return menuList, nil
}

func GetMenuListWithCondition(w map[string]interface{}) ([]*Menu, error) {
	var menuList []*Menu
	where, values, err := WhereBuild(w)
	if err != nil {
		return nil, err
	}
	if err := DB.Self.Where(where, values...).Order("parent_id asc,order_num asc").
		Find(&menuList).Error; err != nil {
		return nil, err
	}
	return menuList, nil
}

func DeleteMenu(id uint64) error {
	var menu Menu
	if err := DB.Self.Where("id = ?", id).Unscoped().Delete(&menu).Error; err != nil {
		return err
	}
	go func() {
		//删除中间表role_menu关联数据
		DeleteRoleMenuByMenuId(id)
	}()
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

//根据用户ID查询菜单
func GetMenuNavByUserId(userId uint64) ([]*Menu, error) {
	var menus []*Menu
	if err := DB.Self.Where("sur.user_id = ?", userId).
		Joins(" left join sys_role_menu srm on srm.menu_id = sys_menu.id" +
			" left join sys_role sr on sr.id = srm.role_id" +
			" left join sys_user_role sur on sur.role_id = sr.id").
		Order("parent_id asc,order_num asc").
		Group("id").Find(&menus).Error;
		err != nil {
		return menus, err
	}
	return menus, nil
}
