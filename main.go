package main

import (
	"time"

	"github.com/HosseinForouzan/E-Commerce-API/config"
	"github.com/HosseinForouzan/E-Commerce-API/delivery"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql/psqluser"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice/authservice"
)

const(
	JWTSignKey = "secret"
	AccessTokenSubject = "at"
	RefreshTokenSubject = "rt"
	AccessTokenExpirationDuration = time.Hour * 24
	RefreshTokenExpirationDuration = time.Hour * 24 * 7

)

func main() {

	cfg := config.Config{
		HttpServer: config.HttpServer{Port: 8080},
		Auth: authservice.Config{
			SignKey: JWTSignKey,
			AccessExpirationTime: AccessTokenExpirationDuration,
			RefreshExpirationTime: RefreshTokenExpirationDuration,
			AccessSubject: AccessTokenSubject,
			RefreshSubject: RefreshTokenSubject,
		},
		Psql: psql.Config{
			Username: "myuser",
			Password: "secret",
			Port: 5431,
			Host: "localhost",
			DBName: "ecommerce_db",
		},

	}



	authSvc, userSvc := setupServices(cfg)

	server := delivery.New(cfg, authSvc, userSvc)
	server.SetRoutes()

}


func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	psql := psql.New(cfg.Psql)
	PsqlUserRepo := psqluser.New(psql)
	userSvc := userservice.New(PsqlUserRepo, authSvc)

	return authSvc, userSvc
}