package registrar

import (
	"log"
	"swap.io-ledger/src/addressSyncer"
	"swap.io-ledger/src/managers/UsersManager"
	"swap.io-ledger/src/networks"
	"swap.io-ledger/src/serviceRegistry"
)

type Registrar struct {
	usersManager  *UsersManager.UsersManager
	addressSyncer *AddressSyncer.AddressSyncer
	networks      *networks.Networks
}
type Config struct {
	UsersManager  *UsersManager.UsersManager
	AddressSyncer *AddressSyncer.AddressSyncer
	Networks      *networks.Networks
}

func InitialiseRegistrar(config Config) *Registrar {
	return &Registrar{
		usersManager:  config.UsersManager,
		addressSyncer: config.AddressSyncer,
		networks:      config.Networks,
	}
}

func Register(reg *serviceRegistry.ServiceRegistry) {
	var usersManager *UsersManager.UsersManager
	err := reg.FetchService(&usersManager)
	if err != nil {
		log.Panicln(err)
	}

	var addressSyncer *AddressSyncer.AddressSyncer
	err = reg.FetchService(&addressSyncer)
	if err != nil {
		log.Panicln(err)
	}

	var networksInstance *networks.Networks
	err = reg.FetchService(&networksInstance)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(
		InitialiseRegistrar(Config{
			UsersManager:  usersManager,
			AddressSyncer: addressSyncer,
			Networks:      networksInstance,
		}),
	)
}

func (*Registrar) Start() {}
func (*Registrar) Stop() error {
	return nil
}
func (*Registrar) Status() error {
	return nil
}
