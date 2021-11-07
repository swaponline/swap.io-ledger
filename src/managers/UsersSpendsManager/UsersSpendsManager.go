package UsersSpendsManager

import (
	"log"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
)

type UsersSpendsManager struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseUsersSpendsManager(config Config) *UsersSpendsManager {
	return &UsersSpendsManager{
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
		InitialiseUsersSpendsManager(Config{
			Database: database,
		}),
	)
}

func (*UsersSpendsManager) Start() {}
func (*UsersSpendsManager) Status() error {
	return nil
}
func (*UsersSpendsManager) Stop() error {
	return nil
}