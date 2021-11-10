package addressSyncer

import (
	"log"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/txsHandler"
)

type AddressSyncer struct {
	database 	  *database.Database
	agentHandlers map[string]*agentHandler.AgentHandler
	txsHandler    *txsHandler.TxsHandler
	onSyncEvents  chan struct{}
}
type Config struct {
	Database 	  *database.Database
	AgentHandlers map[string]*agentHandler.AgentHandler
	TxsHandler    *txsHandler.TxsHandler
	OnSyncEvents  chan struct{}
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

func (*AddressSyncer) Start() {

}
func (*AddressSyncer) Status() error {
	return nil
}
func (*AddressSyncer) Stop() error {
	return nil
}