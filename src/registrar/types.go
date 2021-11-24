package registrar

import "swap.io-ledger/src/database"

type RegistrarData struct {
	PubKey string
	Addresses []database.CreateUserAddressData
}
