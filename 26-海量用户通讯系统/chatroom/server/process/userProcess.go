package processes

import (
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
	"golang_study/26-海量用户通讯系统/chatroom/server/model"
	"golang_study/26-海量用户通讯系统/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

//编写一个函数serverProcessLogin函数， 专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码...
	//1. 先从mes中取出mes.Data ，并直接反序列化成LoginMes
	var loginMes message.Login_Message
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.Login_Res_Mes_Type

	//再声明一个 loginResMes
	var loginResMes message.Login_Response_Message

	//我们需要到redis数据库去验证，用户身份是否合法

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()

		} else {
			loginResMes.Code = 505
			fmt.Println("服务器内部发生错误...")
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登陆成功")
	}
	//3将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6.发送
	//创建一个 Transfer 实例
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	return
}
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.Register_Message
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.Register_Res_Mes_Type

	//再声明一个 loginResMes
	var registerResMes message.Register_Response_Message

	//我们需要到redis数据库去验证，用户身份是否合法

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"

		}
	} else {
		registerResMes.Code = 200
		fmt.Println("用户注册成功")
	}
	//3将loginResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6.发送
	//创建一个 Transfer 实例
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	return
}
