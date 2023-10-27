package gormx

import (
	"cladmin/pkg/errno"
	"errors"
	"github.com/chenglongcl/log"
	"gorm.io/gorm"
)

func HandleError(err error) *errno.Errno {
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("database error", err)
		return errno.ErrDatabase
	}
	return nil
}
