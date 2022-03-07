package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filepath1 := "./a.txt"
	filepath2 := "./b.txt"

	data, err := ioutil.ReadFile(filepath1)
	if err != nil {
		fmt.Printf("reader err=%v\n", err)
	}

	err = ioutil.WriteFile(filepath2, data, 0666)
	if err != nil {
		fmt.Printf("writer err=%v\n", err)
	}
}
