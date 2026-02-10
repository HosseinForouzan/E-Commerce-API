package main

import (

	"github.com/HosseinForouzan/E-Commerce-API/delivery"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql/psqluser"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice"
)

func main() {


	psql := psql.New()
	userpsql := psqluser.New(psql)
	userSVc := userservice.New(userpsql)

	server := delivery.New(userSVc)
	server.SetRoutes()

}