package main

import (
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql/psqluser"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice"
)

func main() {
	psql := psql.New()
	userpsql := psqluser.New(psql)
	fmt.Println(userpsql)

	user := param.RegisterRequest{
		Name: "Hossein",
		Email: "Hossein",
		Password: "Hossein",
		PhoneNumber: "0912",
	}
	userSVc := userservice.New(psql)

	r, err := userSVc.Register(user)
	if err != nil{
		fmt.Errorf("%W", err)
	}

	fmt.Println(r)


}