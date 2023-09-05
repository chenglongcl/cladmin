package sd

import (
	"cladmin/handler"
	"cladmin/service/demoservice"
	"github.com/gin-gonic/gin"
)

func DemoOne(c *gin.Context) {
	demoService := demoservice.NewDemoService(c)
	userModel, errNo := demoService.DemoOne()
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	handler.SendResponse(c, nil, userModel)
}
