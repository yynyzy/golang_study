package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局pool
var pool *redis.Pool

//根据用户配置初始化一个 链接池
func initPool(adress string, MaxIdle, MaxAtive int, IdleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxAtive,
		IdleTimeout: IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", adress)
		},
	}
}
