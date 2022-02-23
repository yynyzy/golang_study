package model

type student struct {
	Name  string
	Score float64
}

//因为 student 结构体首字母是小写，因此只能子啊model使用
//我们通过工厂模式来解决
