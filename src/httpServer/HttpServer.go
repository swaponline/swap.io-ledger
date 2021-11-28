package HttpServer

import (
	"log"
	"swap.io-ledger/src/registrar"
	"swap.io-ledger/src/serviceRegistry"
)

type HttpServer struct {
	registrar *registrar.Registrar
}
type Config struct {
	Registrar *registrar.Registrar
}

func InitialiseHttpServer(config Config) *HttpServer {
	httpServer := &HttpServer{
		registrar: config.Registrar,
	}
	httpServer.handleRegistration()

	return httpServer
}
func Register(reg *serviceRegistry.ServiceRegistry) {
	var registrarInstance *registrar.Registrar
	err := reg.FetchService(&registrarInstance)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(
		 InitialiseHttpServer(Config{
			 Registrar: registrarInstance,
		 }),
	)
	if err != nil {
		log.Panicln(err)
	}
}


func (*HttpServer) Start() {}
func (*HttpServer) Status() error  {
	return nil
}
func (*HttpServer) Stop() error {
	return nil
}
