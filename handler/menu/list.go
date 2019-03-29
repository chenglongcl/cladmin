package menu

import (
	. "cladmin/handler"
	"cladmin/service/menu_service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	menuService := menu_service.Menu{}
	info, errNo := menuService.GetList()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, info)
}
