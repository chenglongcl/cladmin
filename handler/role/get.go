package role

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/role_service"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	roleService := role_service.Role{
		Id: r.Id,
	}
	role, errNo := roleService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	var menuIdList []uint64
	for _, v := range role.Menu {
		menuIdList = append(menuIdList, v.Id)
	}
	SendResponse(c, nil, GetResponse{
		Id:           role.Id,
		RoleName:     role.RoleName,
		Remark:       role.Remark,
		CreateUserId: role.CreateUserId,
		MenuIdList:   menuIdList,
		CreateTime:   role.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
