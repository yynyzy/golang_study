package processes

import (
	"encoding/json"
	"fmt"
	"golang_study/26-海量用户通讯系统/chatroom/common/message"
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

	//如果用户id= 100,密码=123,认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法,返回500
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册后再登陆..."
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
