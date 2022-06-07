package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Person struct {
	XMLName xml.Name `xml : "person"`
	Name    string   `xml : "name"`
	Age     int      `xml : "age"`
	Email   string   `xml : "email"`
}

func Marshal() {
	person := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	//b,_ := xml.Marshal(person)
	b, _ := xml.MarshalIndent(person, " ", " ")
	fmt.Printf("%v\n", string(b))
}
func main() {
	b, _ := ioutil.ReadFile("go_study\\24-golang标准库\\10-xml包\\a.xml")
	var p Person
	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}
