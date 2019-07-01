package configservice

import (
	"cladmin/model"
	"cladmin/pkg/errno"
)

type Config struct {
	ID         uint64
	ParamKey   string
	ParamValue string
	Status     int64
	Type       int64
	Remark     string
}

func (a *Config) GetByParamKey() (*model.Config, *errno.Errno) {
	config, err := model.GetConfigByParamKey(a.ParamKey)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return config, nil
}

func (a *Config) Edit() *errno.Errno {
	data := map[string]interface{}{
		"id":          a.ID,
		"param_key":   a.ParamKey,
		"param_value": a.ParamValue,
		"status":      a.Status,
		"type":        a.Type,
		"remark":      a.Remark,
	}
	if err := model.EditConfig(data); err != nil {
		return errno.ErrDatabase
	}
	return nil
}
