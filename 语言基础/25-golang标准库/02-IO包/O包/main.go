package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	r := strings.NewReader("Yzy Hello")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	/*
		r := strings.NewReader("Yzy Hello")
		buf := make([]byte, 20)
		r.Read(buf)
		fmt.Printf("string(buf):%v\n", string(buf))
	*/
	testCopy()
}
