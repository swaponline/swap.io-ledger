package AddressSyncer

import (
	"log"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/managers/AddressSyncStatusManager"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/txsHandler"
)

type AddressSyncer struct {
	agentHandlers map[string]*AgentHandler.AgentHandler
	txsHandler    *txsHandler.TxsHandler
	addressSyncStatusManager *AddressSyncStatusManager.AddressSyncStatusManager
	onSyncEvents  chan struct{}
}
type Config struct {
	AgentHandlers map[string]*AgentHandler.AgentHandler
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
	var hsdHandler *AgentHandler.AgentHandler
	err := reg.FetchService(&hsdHandler)
	if err != nil {
		log.Panicln(err)
	}

	var addressSyncStatusManager *AddressSyncStatusManager.AddressSyncStatusManager
	err = reg.FetchService(&addressSyncStatusManager)
	if err != nil {
		log.Panicln(err)
	}

	var txsHandlerInstance *txsHandler.TxsHandler
	err = reg.FetchService(&txsHandlerInstance)
	if err != nil {
		log.Panicln(err)
	}

	agentHandlers := make(map[string]*AgentHandler.AgentHandler)
	agentHandlers["Handshake"] = hsdHandler

	err = reg.RegisterService(
		InitialiseAddressSyncer(Config{
			AgentHandlers: agentHandlers,
			TxsHandler: txsHandlerInstance,
			AddressSyncStatusManager: addressSyncStatusManager,
		}),
	)
}

func (a *AddressSyncer) Start() {
	addressesSyncStatuses, err := a.addressSyncStatusManager.GetNotSyncAddresses();
	if err != nil {
		log.Panicln(err)
	}
	for _, addressSyncStatus := range addressesSyncStatuses {
		a.SyncAddress(&addressSyncStatus)
	}
}
func (*AddressSyncer) Status() error {
	return nil
}
func (*AddressSyncer) Stop() error {
	return nil
}