package role

import (
	. "cladmin/handler"
	"cladmin/model"
	"cladmin/pkg/errno"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := util.Validate(&r); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	data := map[string]interface{}{
		"role_name":      r.RoleName,
		"create_user_id": r.CreateUserId,
		"remark":         r.Remark,
		"menu_id_list":   r.MenuIdList,
	}
	if _, err := model.AddRole(data); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
