package main

import "fmt"

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
func main() {
	test()
}
