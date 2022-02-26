package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func (c Customer) GetInfo() string {
	info := fmt.Sprintf("\t%v\t%v\t%v\t%v\t%v\t%v", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return info
}

func NewCustomer(Id int, Name string, Gender string, Age int, Phone string, Email string) Customer {
	return Customer{
		Id:     Id,
		Name:   Name,
		Gender: Gender,
		Age:    Age,
		Phone:  Phone,
		Email:  Email,
	}
}
