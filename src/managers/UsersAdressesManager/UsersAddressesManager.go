package UsersAdressesManager

import "swap.io-ledger/src/database"

type UsersAddressesManager struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseUsersAddress(config Config) *UsersAddressesManager {
	return &UsersAddressesManager{
		database: config.Database,
	}
}