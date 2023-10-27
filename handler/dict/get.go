package dict

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/dictservice"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func GetDictType(c *gin.Context) {
	var (
		r    GetDictTypeRequest
		resp GetDictTypeResponse
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	dictTypeService := dictservice.NewDictTypeService(c)
	dictTypeModel, errNo := dictTypeService.Get([]field.Expr{
		cladminquery.Q.SysDictType.ALL,
	}, append(func() []gen.Condition {
		conditions := make([]gen.Condition, 0)
		if r.ID != 0 {
			conditions = append(conditions, cladminquery.Q.SysDictType.ID.Eq(r.ID))
		}
		if r.DictType != "" {
			conditions = append(conditions, cladminquery.Q.SysDictType.DictType.Eq(r.DictType))
		}
		return conditions
	}(), []gen.Condition{}...))
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if dictTypeModel == nil || dictTypeModel.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	if err := copier.Copy(&resp, dictTypeModel); err != nil {
		handler.SendResponse(c, errno.ErrParams, nil)
		return
	}
	handler.SendResponse(c, nil, resp)
}

func GetDictData(c *gin.Context) {
	var (
		r    GetDictDataRequest
		resp GetDictDataResponse
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	dictDataService := dictservice.NewDictDataService(c)
	dictDataModel, errNo := dictDataService.Get([]field.Expr{
		cladminquery.Q.SysDictData.ALL,
	}, append(func() []gen.Condition {
		conditions := make([]gen.Condition, 0)
		if r.ID != 0 {
			conditions = append(conditions, cladminquery.Q.SysDictData.ID.Eq(r.ID))
		}
		return conditions
	}(), []gen.Condition{}...))
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if dictDataModel == nil || dictDataModel.ID == 0 {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	if err := copier.Copy(&resp, dictDataModel); err != nil {
		handler.SendResponse(c, errno.ErrParams, nil)
		return
	}
	handler.SendResponse(c, nil, resp)
}
