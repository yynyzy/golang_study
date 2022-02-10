package main

import "fmt"

func main() {
	monster := make([]map[string]string, 2)
	if monster == nil {
		monster[0]["name"] = "yyn"
		monster[0]["age"] = "18"

		monster[1]["name"] = "yzy"
		monster[1]["age"] = "17"
	}
	fmt.Println(monster)
}
