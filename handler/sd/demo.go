package sd

import (
	"github.com/gin-gonic/gin"
	"apiserver/model"
	"github.com/json-iterator/go"
)

func DemoOne(c *gin.Context) {
	user, err := model.GetUser(1, []string{"id", "username", "mobile"})
	if err != nil {
		panic(err)
	}
	json, _ := jsoniter.Marshal(user)
	model.RD.Self.Do("HSET", model.RD.Key+"users", user.Id, json)
}
