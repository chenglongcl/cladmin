package model

import "github.com/spf13/viper"

type Config struct {
	BaseModel
	ParamKey   string `gorm:"column:param_key"`
	ParamValue string `gorm:"column:param_value"`
	Status     int64  `gorm:"column:status"`
	Type       int64  `gorm:"column:type"`
	Remark     string `gorm:"column:remark"`
}

func (c *Config) TableName() string {
	return viper.GetString("db.prefix") + "config"
}

func GetConfigByParamKey(str string) (*Config, error) {
	var config Config
	if err := SelectDB("self").Where("param_key = ?", str).First(&config).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

func EditConfig(data map[string]interface{}) error {
	var config Config
	return SelectDB("self").Model(&config).Where("id = ?", data["id"]).Updates(data).Error
}
