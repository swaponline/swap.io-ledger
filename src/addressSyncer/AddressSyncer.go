package addressSyncer

import (
	"log"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
)

type AddressSyncer struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseAddressSyncer(config Config) *AddressSyncer {
	return &AddressSyncer{
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
		InitialiseAddressSyncer(Config{
			Database: database,
		}),
	)
}

func (*AddressSyncer) Start() {}
func (*AddressSyncer) Status() error {
	return nil
}
func (*AddressSyncer) Stop() error {
	return nil
}