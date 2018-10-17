package article

import (
	"github.com/gin-gonic/gin"
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/util"
	"apiserver/service"
	"github.com/lexkong/log"
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
	log.Infof("%+v\n", r)
	ps.Setting(r.Page)
	infos, count, err := service.ListArticle(r.CateId, ps)
	if err != nil {
		SendResponse(c, err, nil)
	}
	SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, infos))
}
