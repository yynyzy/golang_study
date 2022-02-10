package main

import "fmt"

func main() {
	monster := make([]map[string]string, 2)
	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["name"] = "yyn"
		monster[0]["age"] = "18"
	}
	if monster[1] == nil {
		monster[1] = make(map[string]string, 2)
		monster[1]["name"] = "yzy"
		monster[1]["age"] = "17"
	}
	fmt.Println(monster)
}
