package sd

import (
	. "cladmin/handler"
	"cladmin/service/demo_service"
	"github.com/gin-gonic/gin"
)

func DemoOne(c *gin.Context) {
	demoService := &demo_service.Demo{}
	info := demoService.DemoOne()
	SendResponse(c, nil, info)
}
