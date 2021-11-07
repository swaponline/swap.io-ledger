package CoinsManager

import (
	"log"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
)

type CoinsManager struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseCoinsManager(config Config) *CoinsManager {
	return &CoinsManager{
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
		InitialiseCoinsManager(Config{
			Database: database,
		}),
	)
}

func (*CoinsManager) Start() {}
func (*CoinsManager) Status() error {
	return nil
}
func (*CoinsManager) Stop() error {
	return nil
}
