package menu

import (
	. "cladmin/handler"
	"cladmin/service/menuservice"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	menuService := menuservice.Menu{}
	w := map[string]interface{}{}
	info, errNo := menuService.GetList(w)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, info)
}

//上级菜单type 为1,2类型
func Select(c *gin.Context) {
	menuService := menuservice.Menu{}
	w := map[string]interface{}{
		"type !=": 2,
	}
	info, errNo := menuService.GetList(w)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, info)
}
