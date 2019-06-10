package common

import (
	"cladmin/pkg/errno"
	"mime/multipart"
)

type OSSClient interface {
	UpLoad(file multipart.File, header *multipart.FileHeader) (string, *errno.Errno)
	ResetClient() bool
}