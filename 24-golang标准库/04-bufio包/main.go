package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func test1() {
	f, _ := os.Open("go_study\\24-golang标准库\\04-bufio包\\b.txt")
	defer f.Close()
	r2 := bufio.NewReader(f)

	for {
		s, err := r2.ReadString('\n')
		fmt.Println(s)
		if err == io.EOF {
			break
		}
	}
}
func test2() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	p := make([]byte, 10)
	for {
		n, err := br.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p):%v\n", string(p[:n]))
		}

	}
}
func test3() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)
	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	br.UnreadByte() //吐出一个字节
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
}
func main() {
	// test1()
	// test2()
	test3()
}
