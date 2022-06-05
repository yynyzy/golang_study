package main

import (
	"bytes"
	"fmt"
	"strings"
)

func test() {
	var i int = 100
	var b byte = 10
	b = byte(i)

	i = int(b)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}

	s = string(b1)

	b1 = []byte(s)
	fmt.Println(s)
	fmt.Println(b1)
}

func test2() {
	s := "duoke360. com"
	b := []byte(s)
	b1 := []byte("duoke360")
	b2 := []byte("DuoKe360")
	strings.Contains("hello world", "hello")
	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)
	b3 = bytes.Contains(b, b2)
	fmt.Printf("b3: %v\n", b3)

}
func main() {
	// test()
	test2()
}
