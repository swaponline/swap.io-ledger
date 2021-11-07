package TxsManager

import (
	"log"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
)

type TxsManager struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseTxsManager(config Config) *TxsManager {
	return &TxsManager{
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
		InitialiseTxsManager(Config{
			Database: database,
		}),
	)
}

func (*TxsManager) Start() {}
func (*TxsManager) Status() error {
	return nil
}
func (*TxsManager) Stop() error {
	return nil
}
