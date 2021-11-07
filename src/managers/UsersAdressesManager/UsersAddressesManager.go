package UsersAdressesManager

import (
	"log"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
)

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
func Register(reg *serviceRegistry.ServiceRegistry) {
	var database *database.Database
	err := reg.FetchService(&database)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(
		InitialiseUsersAddress(Config{
			Database: database,
		}),
	)
}

func (*UsersAddressesManager) Start() {}
func (*UsersAddressesManager) Status() error {
	return nil
}
func (*UsersAddressesManager) Stop() error {
	return nil
}