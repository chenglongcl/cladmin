package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

type User struct {
	BaseModel
	Username     string `gorm:"column:username;not null,unique_index"`
	Password     string `gorm:"column:password;not null"`
	Mobile       string `gorm:"column:mobile"`
	Email        string `gorm:"column:email"`
	Status       int64  `gorm:"column:status"`
	CreateUserID uint64 `gorm:"column:create_user_id"`
	Role         []Role `gorm:"many2many:sys_user_role;"`
}

type UserInfo struct {
	ID           uint64 `json:"userId"`
	Username     string `json:"username"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
	Status       int64  `json:"status"`
	CreateUserID uint64 `json:"createUserId"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

func (u *User) TableName() string {
	return viper.GetString("db.prefix") + "user"
}

func GetUserByUsername(username string) (*User, error) {
	u := &User{}
	d := SelectDB("self").Where("username = ?", username).First(&u)
	return u, d.Error
}

func CheckUserByID(id uint64) (bool, error) {
	var user User
	err := SelectDB("self").Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func CheckUserByUsername(username string) (bool, error) {
	var user User
	err := SelectDB("self").Select("id").Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func CheckUserUsernameID(username string, id uint64) (bool, error) {
	var user User
	err := SelectDB("self").Where("username = ? AND id != ?", username, id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddUser(data map[string]interface{}) (id uint64, err error) {
	user := User{
		Username:     data["username"].(string),
		Password:     data["password"].(string),
		Mobile:       data["mobile"].(string),
		Email:        data["email"].(string),
		Status:       data["status"].(int64),
		CreateUserID: data["create_user_id"].(uint64),
	}
	var role []Role
	SelectDB("self").Where("id in (?)", data["role_id"].([]int64)).Find(&role)
	if err := SelectDB("self").Create(&user).Association("Role").Append(role).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func EditUser(data map[string]interface{}) error {
	var (
		role []Role
		user User
	)
	SelectDB("self").Where("id in (?)", data["role_id"].([]int64)).Find(&role)
	if err := SelectDB("self").Where("id = ?", data["id"].(uint64)).First(&user).Error; err != nil {
		return err
	}
	SelectDB("self").Model(&user).Association("Role").Replace(role)
	delete(data, "role_id")
	if data["password"].(string) == "" {
		SelectDB("self").Model(&user).Omit("password").Updates(data)
	} else {
		SelectDB("self").Model(&user).Updates(data)
	}
	return nil
}

func EditPersonal(data map[string]interface{}) error {
	var user User
	if err := SelectDB("self").Where("id = ?", data["id"].(uint64)).First(&user).Error; err != nil {
		return err
	}
	if data["password"].(string) == "" {
		SelectDB("self").Model(&user).Omit("password").Updates(data)
	} else {
		SelectDB("self").Model(&user).Updates(data)
	}
	return nil
}

func GetUser(id uint64) (*User, error) {
	var user User
	err := SelectDB("self").Preload("Role").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func GetUserList(w map[string]interface{}, offset, limit uint64) ([]*User, uint64, error) {
	users := make([]*User, 0)
	var count uint64
	where, values, _ := WhereBuild(w)
	if err := SelectDB("self").Model(&User{}).Where(where, values...).Count(&count).Error; err != nil {
		return users, count, err
	}
	if err := SelectDB("self").Model(&User{}).Where(where, values...).Offset(offset).Limit(limit).Order("id desc").
		Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}

func GetUsersAll() ([]*User, error) {
	var user []*User
	err := SelectDB("self").Preload("Role").Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id uint64) error {
	var user User
	SelectDB("self").Where("id = ?", id).Preload("Role").First(&user)
	SelectDB("self").Model(&user).Association("Role").Delete(user.Role)
	if err := SelectDB("self").Unscoped().Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
