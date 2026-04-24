package main

import (
	"time"

	"github.com/HosseinForouzan/E-Commerce-API/config"
	"github.com/HosseinForouzan/E-Commerce-API/delivery"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql"
	psqlaccesscontrol "github.com/HosseinForouzan/E-Commerce-API/repository/psql/psqlAccessControl"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql/psqlcart"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql/psqlproduct"
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql/psqluser"
	"github.com/HosseinForouzan/E-Commerce-API/service/authorizationservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/authservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/cartservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/productservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice"
)

const(
	JWTSignKey = "secret"
	AccessTokenSubject = "at"
	RefreshTokenSubject = "rt"
	AccessTokenExpirationDuration = time.Hour * 24
	RefreshTokenExpirationDuration = time.Hour * 24 * 7

)

func main() {

	cfg := config.Load("config.yml")


	// cfg := config.Config{
	// 	HttpServer: config.HttpServer{Port: 8080},
	// 	Auth: authservice.Config{
	// 		SignKey: JWTSignKey,
	// 		AccessExpirationTime: AccessTokenExpirationDuration,
	// 		RefreshExpirationTime: RefreshTokenExpirationDuration,
	// 		AccessSubject: AccessTokenSubject,
	// 		RefreshSubject: RefreshTokenSubject,
	// 	},
	// 	Psql: psql.Config{
	// 		Username: "myuser",
	// 		Password: "secret",
	// 		Port: 5431,
	// 		Host: "localhost",
	// 		DBName: "ecommerce_db",
	// 	},

	// }



	authSvc, userSvc, productSvc, authorizationSvc, cartSvc := setupServices(cfg)

	server := delivery.New(cfg, authSvc, userSvc, productSvc, authorizationSvc, cartSvc)
	server.SetRoutes()

}


func setupServices(cfg config.Config) (authservice.Service, userservice.Service,
	 productservice.Service, authorizationservice.Service, cartservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	psql := psql.New(cfg.Psql)
	PsqlUserRepo := psqluser.New(psql)
	PsqlProductRepo := psqlproduct.New(psql)
	psqlAccessRepo := psqlaccesscontrol.New(psql)
	psqlCartRepo := psqlcart.New(psql)

	userSvc := userservice.New(PsqlUserRepo, authSvc)
	productSvc := productservice.New(PsqlProductRepo)
	authorizationSvc := authorizationservice.New(psqlAccessRepo)
	cartSvc := cartservice.New(psqlCartRepo, PsqlProductRepo)

	return authSvc, userSvc, productSvc, authorizationSvc, cartSvc
}