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

func (*UsersAddressesManager) Start() {}
func (*UsersAddressesManager) Status() error {
	return nil
}
func (*UsersAddressesManager) Stop() error {
	return nil
}