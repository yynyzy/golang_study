package message

const (
	Login_Mes_Type              = "Login_Message"
	Login_Res_Mes_Type          = "Login_Response_Message"
	Register_Mes_Type           = "Register_Message"
	Register_Res_Mes_Type       = "Register_Response_Message"
	Notify_User_Status_Mes_Type = "Notify_User_Status_Mes"
)

//定义几个用户状态常量
const (
	UserOline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息数据
}

type Login_Message struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

//登陆返回的消息
type Login_Response_Message struct {
	Code    int    `json:"code"` //状态码
	UsersId []int  //保存用户Id的切片
	Error   string `json:"error"` //错误的类型
}

//登陆返回的消息
type Register_Message struct {
	User User `json:"user"` //类型就是User结构体
}

type Register_Response_Message struct {
	Code  int    `json:"code"`  //状态码400表示用户已占有，200表示注册成功
	Error string `json:"error"` //错误的类型
}

//为了配合服务器推送用户状态变化的字段
type Notify_User_Status_Mes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
