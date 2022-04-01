package model

type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，我们必须保证
	//用户信息的 json 字符串的 key 和结构体的字段对应的tag名字一致!!!
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
