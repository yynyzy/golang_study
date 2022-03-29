package message

const (
	Login_Mes_Type     = "Login_Message"
	Login_Res_Mes_Type = "Login_Response_Message"
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
	Code  int   `json:"code"`  //状态码
	Error error `json:"error"` //错误的类型
}
