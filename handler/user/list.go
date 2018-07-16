package user

import (
	"github.com/gin-gonic/gin"
	"apiserver/util"
	. "apiserver/handler"
	"apiserver/service"
	"apiserver/pkg/errno"
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
	infos, count, err := service.ListUser(r.UserName, ps)
	if err != nil {
		SendResponse(c, err, nil)
	}
	SendResponse(c, nil, util.PageUtil(count, r.Page, ps.Limit, infos))
}
