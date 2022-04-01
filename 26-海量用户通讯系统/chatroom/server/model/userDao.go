package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	pool *redis.Pool
}

//1.根据用户id 返回一个 用户实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGET", "user", "id"))
	if err != nil {
		if err == redis.ErrNil {
			//表示在 user 哈希中，没有找到对应的 id
			err = ERROR_USER_NOT_EXISTS
		}
		return
	}
	user = &User{}
	//这里我们需要把 res 反序列化成User 实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}
