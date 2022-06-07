package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//从输出流解析
func test6() {
	f, _ := os.Open("go_study\\19-JSON序列化\\03-Decode\\a.json")
	defer f.Close()
	d := json.NewDecoder(f)
	var v map[string]interface{}
	d.Decode(&v)
	fmt.Printf("v : %v\n", v)
}
func main() {
	test6()
}
