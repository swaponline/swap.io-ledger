package auth

import (
	"log"
	"swap.io-ledger/src/managers/UsersManager"
	"swap.io-ledger/src/serviceRegistry"
)

type Auth struct {
	usersManager *UsersManager.UsersManager
}
type Config struct {
	UsersManager *UsersManager.UsersManager
}

func InitialiseAuth(config Config) *Auth {
	return &Auth{
		usersManager: config.UsersManager,
	}
}
func Register(reg *serviceRegistry.ServiceRegistry) {
	var usersManager *UsersManager.UsersManager
	err := reg.FetchService(&usersManager)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(InitialiseAuth(Config{
		usersManager,
	}))
	if err != nil {
		log.Panicln(err)
	}
}

func (*Auth) Start() {}
func (*Auth) Status() error {
	return nil
}
func (*Auth) Stop() error {
	return nil
}