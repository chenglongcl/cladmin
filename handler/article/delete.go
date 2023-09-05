package article

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/articleservice"
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
			articleService := articleservice.NewArticleService(c)
			if errNo := articleService.Delete([]gen.Condition{
				cladminquery.Q.SysArticle.ID.Eq(id),
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
