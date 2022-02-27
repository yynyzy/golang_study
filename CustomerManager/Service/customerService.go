package service

import (
	model "golang_study/CustomerManager/Model"
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
