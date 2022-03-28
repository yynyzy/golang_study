package message

const (
	Login_Mes_Type     = "Login_Message"
	Login_Res_Mes_Type = "Login_Response_Message"
)

type Message struct {
	Type string //消息类型
	Data string //消息数据
}

type Login_Message struct {
	UserId   int
	UserPwd  string
	UserName string
}

//登陆返回的消息
type Login_Response_Message struct {
	Code  int   //状态码
	Error error //错误的类型
}
