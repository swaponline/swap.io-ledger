package registrar

import (
	"swap.io-ledger/src/managers/UsersAdressesManager"
	"swap.io-ledger/src/managers/UsersManager"
)

type Registrar struct {

}
type Config struct {
	UsersManager *UsersManager.UsersManager
	UsersAddressesManager *UsersAdressesManager.UsersAddressesManager
}

func InitialiseRegistrar(config Config) *Registrar {
	return &Registrar{

	}
}

