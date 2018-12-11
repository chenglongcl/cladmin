package sd

import (
	"apiserver/service"
	"github.com/gin-gonic/gin"
)

func DemoOne(c *gin.Context) {
	service.DemoOne()
}
