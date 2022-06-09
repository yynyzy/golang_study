package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/*
Redis链接池
说明:通过Golang对Redis操作，还可 以通过Redis链接池,流程如下:
	1)事先初始化一定数量的链接，放入到链接池
	2)当Go 需要操作Redis时，直接从Redis链接池取出链接即可。
	3) 这样可以节省临时获取Redis链接的时间，从而提高效率.
*/
/*
var pool *redis.Pool
pool = &redis.Pool{
		Maxldle: 8, //最大空闲链接数
		MaxActive: 0, //表示和数据库的最大链接数, 0表示没有限制
		ldleTimeout:100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			//初始化链接的代码，链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
}
c := pool.Get() //从连接池中取出-个链接
pool.Close()//关闭连接池， 旦关闭链接池，就不能从链接池再职
*/

var pool *redis.Pool

func main() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	//从 pool 中取一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "严致远")
	if err != nil {
		fmt.Println("redis Set err! err=", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis Get err! err=", err)
		return
	}
	fmt.Println("redis Get Name=", r)
}
