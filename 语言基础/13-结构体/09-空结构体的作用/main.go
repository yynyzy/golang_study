package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := struct{}{}
	println(unsafe.Sizeof(a))
	// Output: 0

	set := make(map[string]struct{})
	for _, value := range []string{"apple", "orange", "apple"} {
		set[value] = struct{}{}
	}
	fmt.Println(set)
	// Output: map[orange:{} apple:{}]
}
