package bulletin

import (
	"cladmin/dal/cladmindb/cladminquery"
	"cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/bulletinservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	bulletinService := bulletinservice.NewBulletinService(c)
	updateData := make(map[string]interface{})
	updateData["title"] = r.Title
	updateData["tag"] = r.Tag
	updateData["content"] = r.Content
	if errNo := bulletinService.Edit([]gen.Condition{
		cladminquery.Q.SysBulletin.ID.Eq(r.ID),
	}, updateData); errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
