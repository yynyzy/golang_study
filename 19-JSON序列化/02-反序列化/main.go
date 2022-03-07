package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string
	Age  int
}

func UnmarshalStruct() {
	var monster Monster
	str := "{\"Name\":\"严以诺\",\"Age\":18}"
	data, err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("UnmarshalStruct err=%v\n", err)
	} else {
		fmt.Printf("monster反序列化%v\n", string(data))
	}
}

func TestMap() {
	//定义一个map
	var demo map[string]interface{} = make(map[string]interface{})
	demo["name"] = "阿斯顿"
	demo["age"] = 20
	data, err := json.Marshal(&demo)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	} else {
		fmt.Printf("Map序列化%v\n", string(data))
	}
}

func TestSlice() {
	var sliceArray []map[string]interface{}
	var s1 = make(map[string]interface{})
	s1["name"] = "北京"
	s1["age"] = 99
	var s2 = make(map[string]interface{})
	s2["name"] = "湖南"
	s2["age"] = 88

	sliceArray = append(sliceArray, s1, s2)
	data, err := json.Marshal(&sliceArray)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	} else {
		fmt.Printf("切片序列化%v\n", string(data))
	}
}

func main() {
	UnmarshalStruct()
	UnmarshalMap()
	UnmarshalSlice()
}
