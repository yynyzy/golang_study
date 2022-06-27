package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go 向redis写入数据和读取数据
	//1。链接到redis
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis connect err! err=", err)
		return
	}
	defer c.Close() //关闭...

	//2.通过go 向redis写入数据 string [key-val]
	_, err = c.Do("Set", "Name", "Yanzhiyuan")

	if err != nil {
		fmt.Println("redis Set err! err=", err)
		return
	}

	//3.通过go 向redis读取数据 string[key-val]
	r, err := redis.String(c.Do("Get", "Name"))
	if err != nil {
		fmt.Println("redis Get err! err=", err)
		return
	}

	//返回的r是 interface{}类型,r 是一个 []byte
	//所以对r需要转换，使用redis.String()
	fmt.Println("redis Get Name=", r)
	fmt.Println("redis connect success", c)

}
