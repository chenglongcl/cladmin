package article

import (
	"github.com/gin-gonic/gin"
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/util"
	"apiserver/service"
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
	ps.Setting(r.Page)
	infos, count, err := service.ListArticle(r.CateId, ps)
	if err != nil {
		SendResponse(c, err, nil)
	}
	SendResponse(c, nil, util.PageUtil(count, r.Page, ps.Limit, infos))
}
