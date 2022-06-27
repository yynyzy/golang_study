package main

import "fmt"

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}
func main() {
	var a int = 8
	var b int = 5
	swap(&a, &b)
	fmt.Printf("a的值为%v，b的值为%v", a, b)
}
