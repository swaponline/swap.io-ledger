package addressSyncer

import (
	"log"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/managers/AddressSyncStatusManager"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/txsHandler"
)

type AddressSyncer struct {
	agentHandlers map[string]*agentHandler.AgentHandler
	txsHandler    *txsHandler.TxsHandler
	addressSyncStatusManager *AddressSyncStatusManager.AddressSyncStatusManager
	onSyncEvents  chan struct{}
}
type Config struct {
	AgentHandlers map[string]*agentHandler.AgentHandler
	TxsHandler    *txsHandler.TxsHandler
	AddressSyncStatusManager *AddressSyncStatusManager.AddressSyncStatusManager
	OnSyncEvents  chan struct{}
}

func InitialiseAddressSyncer(config Config) *AddressSyncer {
	return &AddressSyncer{
		agentHandlers: config.AgentHandlers,
		txsHandler: config.TxsHandler,
		addressSyncStatusManager: config.AddressSyncStatusManager,
		onSyncEvents: config.OnSyncEvents,
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
			// todo: add managers
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