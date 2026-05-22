package config

import (
	"gameapp/repository/mysql"
	"gameapp/service/authservice"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	HttpServer HTTPServer
	Auth       authservice.Config
	MySQL      mysql.Config
}
