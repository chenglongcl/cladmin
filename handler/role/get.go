package role

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/roleservice"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	roleService := roleservice.NewRoleService(c)
	roleModel, errNo := roleService.Get([]field.Expr{
		cladminquery.Q.SysRole.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysRole.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if roleModel == nil || roleModel.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	menuIDList := make([]uint64, 0)
	_ = jsoniter.UnmarshalFromString(roleModel.MenuIDList, &menuIDList)
	handler.SendResponse(c, nil, GetResponse{
		ID:           roleModel.ID,
		RoleName:     roleModel.RoleName,
		Remark:       roleModel.Remark,
		CreateUserID: roleModel.CreateUserID,
		MenuIDList:   menuIDList,
		CreateTime:   roleModel.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
