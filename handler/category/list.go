package category

import (
	. "cladmin/handler"
	"cladmin/service/category_service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	categoryService := category_service.Category{}
	w := map[string]interface{}{}
	info, errNo := categoryService.GetList(w)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, info)
}
