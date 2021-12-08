package AddressSyncer

import (
	"log"
	"swap.io-ledger/src/managers/AddressSyncStatusManager"
	"swap.io-ledger/src/networks"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/txsHandler"
)

type AddressSyncer struct {
	networks                 *networks.Networks
	txsHandler               *txsHandler.TxsHandler
	addressSyncStatusManager *AddressSyncStatusManager.AddressSyncStatusManager
	onSyncEvents             chan struct{}
}
type Config struct {
	Networks                 *networks.Networks
	TxsHandler               *txsHandler.TxsHandler
	AddressSyncStatusManager *AddressSyncStatusManager.AddressSyncStatusManager
	OnSyncEvents             chan struct{}
}

func InitialiseAddressSyncer(config Config) *AddressSyncer {
	return &AddressSyncer{
		networks:                 config.Networks,
		txsHandler:               config.TxsHandler,
		addressSyncStatusManager: config.AddressSyncStatusManager,
		onSyncEvents:             config.OnSyncEvents,
	}
}
func Register(reg *serviceRegistry.ServiceRegistry) {
	var networksInstance *networks.Networks
	err := reg.FetchService(&networksInstance)
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

	err = reg.RegisterService(
		InitialiseAddressSyncer(Config{
			Networks:                 networksInstance,
			TxsHandler:               txsHandlerInstance,
			AddressSyncStatusManager: addressSyncStatusManager,
		}),
	)
}

func (a *AddressSyncer) Start() {
	addressesSyncStatuses, err := a.addressSyncStatusManager.GetNotSyncAddresses()
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
