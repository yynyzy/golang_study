package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.145.1:8888")
	if err != nil {
		fmt.Printf("client Dial err=%v\n", err)
		return
	}
	fmt.Println("Connecting 成功=", conn)
}
