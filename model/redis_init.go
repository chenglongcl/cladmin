package model

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

type Redis struct {
	Self redis.Conn
	Key  string
}

var RD *Redis

func (rd *Redis) Init() {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(viper.GetString("redis_conf.network"),
				viper.GetString("redis_conf.address"),
				redis.DialPassword(viper.GetString("redis_conf.password")))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	RD = &Redis{
		Self: pool.Get(),
		Key:  viper.GetString("redis_conf.prefix"),
	}
}
func (rd *Redis) Close() {
	RD.Self.Close()
}
