package config

import (
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice/authservice"
)

type HttpServer struct {
	Port int
}

type Config struct {
	HttpServer HttpServer
	Auth authservice.Config
	Psql psql.Config
}