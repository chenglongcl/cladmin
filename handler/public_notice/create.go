package public_notice

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/public_notice_service"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	publicNoticeService := &public_notice_service.PublicNotice{
		Title:   r.Title,
		Tag:     r.Tag,
		Content: r.Content,
	}
	if errNo := publicNoticeService.Add(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}