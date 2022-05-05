package main

import "fmt"

func main() {
	exp := "ch+"
	exp2 := "c"
	exp1 := exp[0:1]
	fmt.Printf("%T,%s\n", exp, exp)
	fmt.Printf("%T,%s\n", exp2, exp2)
	fmt.Printf("%T,%s\n", exp1, exp1)
	fmt.Println([]byte(exp1))
}
