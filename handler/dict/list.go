package dict

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service"
	"cladmin/service/dictservice"
	"cladmin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func ListDictType(c *gin.Context) {
	var (
		r  ListDictTypeRequest
		ps util.PageSetting
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ps.Setting(r.Page, r.Limit)
	dictTypeService := dictservice.NewDictTypeService(c)
	infos, count, errNo := dictTypeService.InfoList(&service.ListParams{
		PS: ps,
		Fields: []field.Expr{
			cladminquery.Q.SysDictType.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.DictName != "" {
				conditions = append(conditions, cladminquery.Q.SysDictType.DictName.Like(util.StringBuilder("%", r.DictName, "%")))
			}
			if r.DictType != "" {
				conditions = append(conditions, cladminquery.Q.SysDictType.DictType.Like(util.StringBuilder("%", r.DictType, "%")))
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysDictType.Sort,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, infos))
}

func AllDictType(c *gin.Context) {
	dictTypeService := dictservice.NewDictTypeService(c)
	infos, _, errNo := dictTypeService.InfoList(&service.ListParams{
		PS: util.PageSetting{},
		Options: struct {
			WithoutCount  bool
			Scenes        string
			CustomDBOrder string
			CustomFunc    func() interface{}
		}{WithoutCount: true, Scenes: "all"},
		Fields: []field.Expr{cladminquery.Q.SysDictType.ALL},
		Orders: []field.Expr{
			cladminquery.Q.SysDictType.Sort,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, infos)
}

func ListDictData(c *gin.Context) {
	var (
		r  ListDictDataRequest
		ps util.PageSetting
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ps.Setting(r.Page, r.Limit)
	dictDataService := dictservice.NewDictDataService(c)
	infos, count, errNo := dictDataService.InfoList(&service.ListParams{
		PS: func() util.PageSetting {
			if r.DictTypeID == 0 {
				return ps
			} else {
				return util.PageSetting{}
			}
		}(),
		Fields: []field.Expr{
			cladminquery.Q.SysDictData.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.DictLabel != "" {
				conditions = append(conditions, cladminquery.Q.SysDictData.DictLabel.Like(util.StringBuilder("%", r.DictLabel, "%")))
			}
			if r.DictValue != "" {
				conditions = append(conditions, cladminquery.Q.SysDictData.DictValue.Like(util.StringBuilder("%", r.DictValue, "%")))
			}
			if r.DictTypeID != 0 {
				conditions = append(conditions, cladminquery.Q.SysDictData.DictTypeID.Eq(r.DictTypeID))
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			cladminquery.Q.SysDictData.Sort,
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, infos))
}
