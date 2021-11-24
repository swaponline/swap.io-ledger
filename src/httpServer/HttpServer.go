package Httpserver

import "swap.io-ledger/src/registrar"

type HttpServer struct {
	registrar *registrar.Registrar
}
type Config struct {
	Registrar *registrar.Registrar
}

func InitialiseHttpServer(config Config) *HttpServer {
	return &HttpServer{
		registrar: config.Registrar,
	}
}
