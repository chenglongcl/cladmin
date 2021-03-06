package bulletin

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/bulletinservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	publicNoticeService := &bulletinservice.Bulletin{
		ID:      r.ID,
		Title:   r.Title,
		Tag:     r.Tag,
		Content: r.Content,
	}
	if errNo := publicNoticeService.Edit(); errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, nil)
}
