package category

import (
	. "cladmin/handler"
	"cladmin/service/categoryservice"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	categoryService := categoryservice.Category{}
	w := map[string]interface{}{}
	info, errNo := categoryService.GetList(w)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, info)
}
