package model

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

type Redis struct {
	Client *redis.Pool
	Key    string
}

var RD *Redis

func (rd *Redis) Init() {
	RD = &Redis{
		Client: &redis.Pool{
			MaxIdle:   16,
			MaxActive: 500,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial(viper.GetString("redis_conf.network"),
					viper.GetString("redis_conf.address"),
					redis.DialPassword(viper.GetString("redis_conf.password")))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		},
		Key: viper.GetString("redis_conf.prefix"),
	}
}
