package dict

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/dictservice"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gen"
)

func DeleteDictType(c *gin.Context) {
	var (
		r DeleteDictTypeRequest
	)
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	errGroup := &errgroup.Group{}
	for _, _id := range r.Ids {
		id := _id
		errGroup.Go(func() error {
			dictTypeService := dictservice.NewDictTypeService(c)
			if errNo := dictTypeService.Delete([]gen.Condition{
				cladminquery.Q.SysDictType.ID.Eq(id),
			}); errNo != nil {
				return errNo
			}
			return nil
		})
	}
	if errNo := errGroup.Wait(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func DeleteDictData(c *gin.Context) {
	var (
		r DeleteDictDataRequest
	)
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	errGroup := &errgroup.Group{}
	for _, _id := range r.Ids {
		id := _id
		errGroup.Go(func() error {
			dictDataService := dictservice.NewDictDataService(c)
			if errNo := dictDataService.Delete([]gen.Condition{
				cladminquery.Q.SysDictData.ID.Eq(id),
			}); errNo != nil {
				return errNo
			}
			return nil
		})
	}
	if errNo := errGroup.Wait(); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
