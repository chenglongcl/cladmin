package redsync

import (
	"cladmin/pkg/redisgo"
	"github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/redigo"
)

var pool redis.Pool

func Init() {
	pool = redigo.NewPool(redisgo.My().GetPool())
}

func GetPool() redis.Pool {
	return pool
}
