package service

import (
	model "golang_study/练习项目-CustomerManager/Model"
)

type CustomerService struct {
	//维护所有客户的切片
	customers []model.Customer
	//表示当前有多少个客户
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "严致远", "女", 18, "13813813812", "1601530253@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}
func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}

	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}
func (cs *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
		}
	}
	return index
}
