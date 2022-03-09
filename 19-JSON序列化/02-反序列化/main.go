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
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("UnmarshalStruct err=%v\n", err)
	} else {
		fmt.Printf("monster反序列化%v\n", monster)
	}
}

func UnmarshalMap() {
	//定义一个map
	var demo map[string]interface{}
	str := "{\"Name\":\"嘤嘤嘤\",\"Age\":20}"
	err := json.Unmarshal([]byte(str), &demo)
	if err != nil {
		fmt.Printf("UnmarshalMap err=%v\n", err)
	} else {
		fmt.Printf("Map反序列化%v\n", demo)
	}
}

func UnmarshalSlice() {
	var sliceArray []map[string]interface{}
	str := "[{\"Name\":\"啧啧啧\",\"Age\": 49},{\"Name\":\"啊啊啊\",\"Age\": 88}]"
	err := json.Unmarshal([]byte(str), &sliceArray)
	if err != nil {
		fmt.Printf("UnmarshalSlice err=%v\n", err)
	} else {
		fmt.Printf("切片反序列化%v\n", sliceArray)
	}
}

func main() {
	UnmarshalStruct()
	UnmarshalMap()
	UnmarshalSlice()
}
