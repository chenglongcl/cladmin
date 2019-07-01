package role

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/roleservice"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	roleService := roleservice.Role{
		ID: r.ID,
	}
	role, errNo := roleService.Get()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	menuIdList := make([]int64, 0)
	jsoniter.UnmarshalFromString(role.MenuIDList, &menuIdList)
	SendResponse(c, nil, GetResponse{
		ID:           role.ID,
		RoleName:     role.RoleName,
		Remark:       role.Remark,
		CreateUserID: role.CreateUserID,
		MenuIDList:   menuIdList,
		CreateTime:   role.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
