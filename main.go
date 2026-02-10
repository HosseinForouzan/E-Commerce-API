package main

import (
	"time"

	"github.com/HosseinForouzan/E-Commerce-API/delivery"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql/psqluser"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice/authservice"
)

func main() {


	psql := psql.New()
	userpsql := psqluser.New(psql)
	authSvc := authservice.New("secret", "at", "rt", time.Hour * 24, time.Hour * 24 * 7)
	userSVc := userservice.New(userpsql, authSvc)

	server := delivery.New(userSVc)
	server.SetRoutes()

}