package bulletin

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/bulletinservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var (
		r  ListRequest
		ps util.PageSetting
	)
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	ps.Setting(r.Page, r.Limit)
	publicNoticeService := bulletinservice.Bulletin{
		Title: r.Title,
		Tag:   r.Tag,
	}

	info, count, errNo := publicNoticeService.GetList(ps)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, info))
}
