package bulletin

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/bulletinservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	publicNoticeService := &bulletinservice.Bulletin{
		Id: r.Id,
	}
	publicNotice, errNo := publicNoticeService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, GetResponse{
		Id:         publicNotice.Id,
		Title:      publicNotice.Title,
		Tag:        publicNotice.Tag,
		Content:    publicNotice.Content,
		CreateTime: publicNotice.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
