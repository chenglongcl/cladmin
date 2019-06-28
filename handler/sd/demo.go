package sd

import (
	. "cladmin/handler"
	"cladmin/service/demoservice"
	"github.com/gin-gonic/gin"
)

func DemoOne(c *gin.Context) {
	demoService := &demoservice.Demo{}
	info := demoService.DemoOne()
	SendResponse(c, nil, info)
}
