package demo_service

import (
	"cladmin/model"
	"cladmin/pkg/redisgo"
)

type Demo struct {
}

func (a *Demo) DemoOne() *model.User {
	user, err := model.GetUser(1)
	if err != nil {
		panic(err)
	}
	redisgo.My().HSet("testUsers", "1", user)
	userTwo := &model.User{}
	redisgo.My().HGetObject("testUsers", "1", userTwo)
	return userTwo
}
