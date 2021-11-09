package AddressSyncStatusManager

import (
	"log"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
)

type AddressSyncStatusManager struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseAddressSyncStatusManager(
	config Config,
) *AddressSyncStatusManager {
	return &AddressSyncStatusManager{
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
		InitialiseAddressSyncStatusManager(Config{
			Database: database,
		}),
	)
}

func (*AddressSyncStatusManager) Start() {}
func (*AddressSyncStatusManager) Status() error {
	return nil
}
func (*AddressSyncStatusManager) Stop() error {
	return nil
}