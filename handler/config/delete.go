package config

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/configservice"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gen"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	errGroup := &errgroup.Group{}
	for _, _id := range r.Ids {
		id := _id
		errGroup.Go(func() error {
			configService := configservice.NewConfigService(c)
			if errNo := configService.Delete([]gen.Condition{
				cladminquery.Q.SysConfig.ID.Eq(id),
				cladminquery.Q.SysConfig.Locked.Is(false),
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
