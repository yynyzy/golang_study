package model

import (
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//此函数用于创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return

}

//1.根据用户id 返回一个 用户实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	res, err := redis.String(conn.Do("HGet", "users", id))

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
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//从链接池中取一个链接
	conn := this.pool.Get()

	defer conn.Close() //关闭这个连接
	//用此链接调用函数判断用户是否存在
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if userPwd != user.UserPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
func (this *UserDao) Register(user *message.User) (err error) {
	//从链接池中取一个链接
	conn := this.pool.Get()

	defer conn.Close() //关闭这个连接
	//判断 redis 中 id 是否存在
	_, err = this.getUserById(conn, user.UserId)
	//如果没有错误，说明redis中存在id，说明用户id已经被注册过
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	data, err := json.Marshal(user) //将传入的注册信息序列化
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("用户注册保存发生错误 err=", err)
		return
	}
	return
}
