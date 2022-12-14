package main

import (
	"fmt"
	"study/patterns/structural/proxy/pkg"
)

var (
	admin = "administrator"
	user  = "user"
	users = map[string]bool{
		"administrator": true,
		"user":          false,
	}
)

func main() {
	proxy := pkg.ProxyDataBase{
		Users: users,
		Db:    &pkg.DataBase{},
	}
	adminData, err := proxy.GetData(admin)
	fmt.Printf("From [%s] Data:[%v] Err:[%v]\n", admin, adminData, err)

	userData, err := proxy.GetData(user)
	fmt.Printf("From [%s] Data:[%v] Err:[%v]\n", user, userData, err)
}
