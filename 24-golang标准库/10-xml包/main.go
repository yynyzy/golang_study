package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml : "person"`
	Name    string   `xml : "name"`
	Age     int      `xml : "age"`
	Email   string   `xml : "email"`
}

func main() {
	person := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	//b,_ := xml.Marshal(person)
	b, _ := xml.MarshalIndent(person, " ", " ")
	fmt.Printf("%v\n", string(b))
}
