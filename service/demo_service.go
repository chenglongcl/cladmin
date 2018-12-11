package service

import (
	"apiserver/model"
	"github.com/json-iterator/go"
)

func DemoOne() {
	redisClient := model.RD.Client.Get()
	defer redisClient.Close()
	user, err := model.GetUser(1, []string{"id", "username", "mobile"})
	if err != nil {
		panic(err)
	}
	json, _ := jsoniter.Marshal(user)
	redisClient.Do("HSET", model.RD.Key+"users", 1, json)
}
