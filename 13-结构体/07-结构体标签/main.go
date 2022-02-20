package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json: "name"`
	Age   int    `json: "age"`
	Skill string `json: "name"`
}

func main() {

	//1．创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}

	//2．将monster变量序列化为json格式字串
	// json.Marshal中使用反射
	jsonstr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonstr", string(jsonstr))
}
