package bulletin

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/bulletinservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func Get(c *gin.Context) {
	var r GetRequest
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	bulletinService := bulletinservice.NewBulletinService(c)
	bulletinModel, errNo := bulletinService.Get([]field.Expr{
		cladminquery.Q.SysBulletin.ALL,
	}, []gen.Condition{
		cladminquery.Q.SysBulletin.ID.Eq(r.ID),
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, GetResponse{
		ID:         bulletinModel.ID,
		Title:      bulletinModel.Title,
		Tag:        bulletinModel.Tag,
		Content:    bulletinModel.Content,
		CreateTime: bulletinModel.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
