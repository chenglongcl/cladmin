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

func CreateDictType(c *gin.Context) {
	var (
		r CreateDictTypeRequest
	)
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	dictTypeService := dictservice.NewDictTypeService(c)
	//查询DictType是否存在
	dictTypeModel, errNo := dictTypeService.Get([]field.Expr{
		cladminquery.Q.SysDictType.ID,
	}, []gen.Condition{
		cladminquery.Q.SysDictType.DictType.Eq(r.DictType),
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
	dictTypeService.DictType = r.DictType
	dictTypeService.DictName = r.DictName
	dictTypeService.Remark = r.Remark
	dictTypeService.Sort = r.Sort
	if _, errNo := dictTypeService.Add(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func CreateDictData(c *gin.Context) {
	var (
		r CreateDictDataRequest
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
	dictDataService.DictTypeID = r.DictTypeID
	dictDataService.DictLabel = r.DictLabel
	dictDataService.DictValue = r.DictValue
	dictDataService.Remark = r.Remark
	dictDataService.Sort = r.Sort
	if _, errNo = dictDataService.Add(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
