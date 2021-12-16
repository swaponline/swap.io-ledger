package HttpServer

import (
	"log"
	"swap.io-ledger/src/auth"
	"swap.io-ledger/src/managers/TxsManager"
	"swap.io-ledger/src/managers/UsersSpendsManager"
	"swap.io-ledger/src/networks"
	"swap.io-ledger/src/registrar"
	"swap.io-ledger/src/serviceRegistry"
)

type HttpServer struct {
	auth               *auth.Auth
	registrar          *registrar.Registrar
	networks           *networks.Networks
	usersSpendsManager *UsersSpendsManager.UsersSpendsManager
	txsManager         *TxsManager.TxsManager
}
type Config struct {
	Auth               *auth.Auth
	Registrar          *registrar.Registrar
	Networks           *networks.Networks
	UsersSpendsManager *UsersSpendsManager.UsersSpendsManager
	TxsManager         *TxsManager.TxsManager
}

func InitialiseHttpServer(config Config) *HttpServer {
	httpServer := &HttpServer{
		auth:               config.Auth,
		registrar:          config.Registrar,
		networks:           config.Networks,
		usersSpendsManager: config.UsersSpendsManager,
		txsManager:         config.TxsManager,
	}
	httpServer.handleRegistration()
	httpServer.handleGetBalances()
	httpServer.handleGetTxs()
	httpServer.handlePushTx()

	return httpServer
}
func Register(reg *serviceRegistry.ServiceRegistry) {
	var authInstance *auth.Auth
	err := reg.FetchService(&authInstance)
	if err != nil {
		log.Panicln(err)
	}

	var registrarInstance *registrar.Registrar
	err = reg.FetchService(&registrarInstance)
	if err != nil {
		log.Panicln(err)
	}

	var networksInstance *networks.Networks
	err = reg.FetchService(&networksInstance)
	if err != nil {
		log.Panicln(err)
	}

	var userSpendsManager *UsersSpendsManager.UsersSpendsManager
	err = reg.FetchService(&userSpendsManager)
	if err != nil {
		log.Panicln(err)
	}

	var txsManager *TxsManager.TxsManager
	err = reg.FetchService(&txsManager)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(
		InitialiseHttpServer(Config{
			Auth:               authInstance,
			Registrar:          registrarInstance,
			Networks:           networksInstance,
			UsersSpendsManager: userSpendsManager,
			TxsManager:         txsManager,
		}),
	)
	if err != nil {
		log.Panicln(err)
	}
}

func (*HttpServer) Start() {}
func (*HttpServer) Status() error {
	return nil
}
func (*HttpServer) Stop() error {
	return nil
}
