package role

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/roleservice"
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
	roleService := roleservice.Role{
		RoleName: r.RoleName,
	}
	roles, count, errNo := roleService.GetList(ps)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, roles))
}

func Select(c *gin.Context) {
	roleService := roleservice.Role{}
	roles, errNo := roleService.GetAll()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, roles)
}
