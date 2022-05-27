package main

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name         int      `json:Name`
	Publications []string `json:Publication,omitempty`
}

func main() {
	t := reflect.TypeOf(Author{})
	fmt.Println("t=", t)
	fmt.Println("t.NumField()=", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		s, _ := t.FieldByName(name)
		fmt.Println(s.Tag)
	}
}
