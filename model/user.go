package model

import (
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"apiserver/pkg/auth"
	"apiserver/pkg/constvar"
	"fmt"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=1,max=128"`
	Mobile   string `json:"mobile" gorm:"column:mobile" validate:"numeric,max=11"`
}

func (u *UserModel) TableName() string {
	return viper.GetString("db.prefix") + "users"
}

func (u *UserModel) CreateUser() error {
	return DB.Self.Create(&u).Error
}
func (u *UserModel) UpdateUser() error {
	return DB.Self.Omit("created_at").Save(u).Error
}
func DeleteUser(userId uint64) error {
	user := &UserModel{}
	user.Id = userId
	return DB.Self.Delete(&user).Error
}
func GetUser(id uint64, fields []string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Select(fields).Where("id = ?", id).First(&u)
	return u, d.Error
}
func GetUserByUsername(username string, fields []string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Select(fields).Where("username = ?", username).First(&u)
	return u, d.Error
}
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
func ListUser(username string, offset uint64, limit uint64) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	users := make([]*UserModel, 0)
	var count uint64
	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}
	if err := DB.Self.Model(&UserModel{}).Where(where).Offset(offset).Limit(limit).Order("id desc").
		Find(&users).Error; err != nil {
		return users, count, err
	}
	return users, count, nil
}
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
