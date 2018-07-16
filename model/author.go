package model

import "github.com/spf13/viper"

type AuthorModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null"`
}

func (u *AuthorModel) TableName() string {
	return viper.GetString("db.prefix") + "users"
}
