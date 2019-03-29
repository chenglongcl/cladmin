package menu

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/menu_service"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	menuService := menu_service.Menu{
		Id: r.Id,
	}
	roleList, errNo := menuService.Delete()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	for _, v := range roleList {
		inject.Obj.Common.RoleAPI.LoadPolicy(v)
	}
	SendResponse(c, nil, nil)
}
