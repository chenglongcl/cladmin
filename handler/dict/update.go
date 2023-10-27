package dict

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/dictservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func UpdateDictType(c *gin.Context) {
	var (
		r UpdateDictTypeRequest
	)
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	dictTypeService := dictservice.NewDictTypeService(c)
	//查询新DictType是否存在
	dictTypeModel, errNo := dictTypeService.Get([]field.Expr{
		cladminquery.Q.SysDictType.ID,
	}, []gen.Condition{
		cladminquery.Q.SysDictType.DictType.Eq(r.DictType),
		cladminquery.Q.SysDictType.ID.Neq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if dictTypeModel != nil && dictTypeModel.ID != 0 {
		handler.SendResponse(c, &errno.Errno{
			Code:    29999,
			Message: "字典类型已存在",
		}, nil)
		return
	}
	//
	updateData := make(map[string]interface{})
	updateData["dict_type"] = r.DictType
	updateData["dict_name"] = r.DictName
	updateData["remark"] = r.Remark
	updateData["sort"] = r.Sort
	if errNo := dictTypeService.Edit([]gen.Condition{
		cladminquery.Q.SysDictType.ID.Eq(r.ID),
	}, updateData); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func UpdateDictData(c *gin.Context) {
	var (
		r UpdateDictDataRequest
	)
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	//查询DictType是否存在
	dictTypeService := dictservice.NewDictTypeService(c)
	dictTypeModel, errNo := dictTypeService.Get([]field.Expr{
		cladminquery.Q.SysDictType.ID,
	}, []gen.Condition{
		cladminquery.Q.SysDictType.ID.Eq(r.DictTypeID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if dictTypeModel == nil || dictTypeModel.ID == 0 {
		handler.SendResponse(c, &errno.Errno{
			Code:    29999,
			Message: "字典类型不存在",
		}, nil)
		return
	}
	//
	dictDataService := dictservice.NewDictDataService(c)
	//查询DictValue是否存在
	dictDataModel, errNo := dictDataService.Get([]field.Expr{
		cladminquery.Q.SysDictData.ID,
	}, []gen.Condition{
		cladminquery.Q.SysDictData.DictTypeID.Eq(r.DictTypeID),
		cladminquery.Q.SysDictData.DictValue.Eq(r.DictValue),
		cladminquery.Q.SysDictData.ID.Neq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	if dictDataModel != nil && dictDataModel.ID != 0 {
		handler.SendResponse(c, &errno.Errno{
			Code:    29999,
			Message: "字典值已存在",
		}, nil)
		return
	}
	updateData := make(map[string]interface{})
	updateData["dict_label"] = r.DictLabel
	updateData["dict_value"] = r.DictValue
	updateData["remark"] = r.Remark
	updateData["sort"] = r.Sort
	if errNo := dictDataService.Edit([]gen.Condition{
		cladminquery.Q.SysDictData.ID.Eq(r.ID),
	}, updateData); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
