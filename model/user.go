package model

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

type User struct {
	BaseModel
	Username     string `json:"username" gorm:"column:username;not null,unique_index"`
	Password     string `json:"password" gorm:"column:password;not null"`
	Mobile       string `json:"mobile" gorm:"column:mobile"`
	Email        string `json:"email" gorm:"column:email"`
	Status       int64  `json:"status" gorm:"column:status"`
	CreateUserId uint64 `json:"create_user_id" gorm:"column:create_user_id"`
	Role         []Role `json:"role" gorm:"many2many:sys_user_role;"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Status    int64  `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

func (u *User) TableName() string {
	return viper.GetString("db.prefix") + "user"
}

func GetUserByUsername(username string, fields []string) (*User, error) {
	u := &User{}
	d := DB.Self.Select(fields).Where("username = ?", username).First(&u)
	return u, d.Error
}

func CheckUserById(id uint64) (bool, error) {
	var user User
	err := DB.Self.Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.Id > 0 {
		return true, nil
	}
	return false, nil
}

func CheckUserByUsername(username string) (bool, error) {
	var user User
	err := DB.Self.Select("id").Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.Id > 0 {
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
		CreateUserId: data["create_user_id"].(uint64),
	}
	var role []Role
	DB.Self.Where("id in (?)", data["role_id"].([]int64)).Find(&role)
	if err := DB.Self.Create(&user).Association("Role").Append(role).Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}

func EditUser(data map[string]interface{}) error {
	var (
		role []Role
		user User
	)
	DB.Self.Where("id in (?)", data["role_id"].([]int64)).Find(&role)
	if err := DB.Self.Where("id = ?", data["id"].(uint64)).First(&user).Error; err != nil {
		return err
	}
	DB.Self.Model(&user).Association("Role").Replace(role)
	delete(data, "role_id")
	if data["password"].(string) == "" {
		DB.Self.Model(&user).Omit("password").Update(data)
	} else {
		DB.Self.Model(&user).Update(data)
	}
	return nil
}

func GetUser(id uint64) (*User, error) {
	var user User
	err := DB.Self.Preload("Role").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func GetUserList(w map[string]interface{}, offset, limit uint64) ([]*User, uint64, error) {
	users := make([]*User, 0)
	var count uint64
	where, values, _ := WhereBuild(w)
	if err := DB.Self.Model(&User{}).Where(where, values...).Count(&count).Error; err != nil {
		return users, count, err
	}
	if err := DB.Self.Model(&User{}).Where(where, values...).Offset(offset).Limit(limit).Order("id desc").
		Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}

func GetUsersAll() ([]*User, error) {
	var user []*User
	err := DB.Self.Preload("Role").Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id uint64) error {
	var user User
	DB.Self.Where("id = ?", id).Preload("Role").First(&user)
	DB.Self.Model(&user).Association("Role").Delete(user.Role)
	if err := DB.Self.Unscoped().Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
