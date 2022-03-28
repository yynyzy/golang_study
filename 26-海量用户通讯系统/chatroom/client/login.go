package main

import "fmt"

func login(UserId int, UserPwd string) (err error) {
	fmt.Printf("%d %s", UserId, UserPwd)
	return
}
