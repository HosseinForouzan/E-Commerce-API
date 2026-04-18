package config

import (
	"github.com/HosseinForouzan/E-Commerce-API/repository/psql"
	"github.com/HosseinForouzan/E-Commerce-API/service/authservice"
)

type HttpServer struct {
	Port int	`koanf:"port"`
}

type Config struct {
	HttpServer HttpServer	`koanf:"http_server"`
	Auth authservice.Config	`koanf:"auth"`
	Psql psql.Config	`koanf:"psql"`
}