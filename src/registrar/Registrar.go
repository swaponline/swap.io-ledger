package registrar

import (
	"swap.io-ledger/src/addressSyncer"
	"swap.io-ledger/src/managers/UsersManager"
)

type Registrar struct {
	usersManager *UsersManager.UsersManager
	addressSyncer *AddressSyncer.AddressSyncer
}
type Config struct {
	UsersManager *UsersManager.UsersManager
	AddressSyncer *AddressSyncer.AddressSyncer
}

func InitialiseRegistrar(config Config) *Registrar {
	return &Registrar{
		usersManager: config.UsersManager,
		addressSyncer: config.AddressSyncer,
	}
}

