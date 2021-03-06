package AddressSyncer

import (
	"log"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/database"
	"time"
)

func (as *AddressSyncer) SyncNewAddresses(
	newAddressesIds []int,
) {
	for _, id := range newAddressesIds {
		syncStatus, err := as.addressSyncStatusManager.GetByAddressId(id)
		if err != nil {
			log.Println(err)
			continue
		}
		if syncStatus.Sync == 0 {
			as.SyncAddress(syncStatus)
		}
	}
}
func (as *AddressSyncer) SyncAddress(
	status *database.AddressSyncStatus,
) {
	if agentHandlerInstance, ok := (*as.networks)[status.Network]; ok {
		log.Println(
			"start sync",
			status.AddressId,
			status.Address,
			status.Network,
		)
		for {
			var cursorTxs *AgentHandler.CursorTxs
			var err error
			if status.Cursor == "null" {
				cursorTxs, err = agentHandlerInstance.GetFirstTxsCursor(status.Address)
			} else {
				cursorTxs, err = agentHandlerInstance.GetTxsCursor(status.Cursor)
			}
			if err != nil {
				log.Println("ERROR SYNC:", err)
				<-time.After(time.Second)
				continue
			}

			log.Println(cursorTxs.Cursor, cursorTxs.NextCursor)
			for _, aTx := range cursorTxs.Transactions {
				as.txsHandler.TxHandle(aTx)
			}

			if cursorTxs.NextCursor != "null" {
				err := as.addressSyncStatusManager.UpdateCursor(
					status.AddressId,
					cursorTxs.NextCursor,
				)
				if err != nil {
					log.Println("ERROR SYNC:", err)
				}
			} else {
				err := as.addressSyncStatusManager.SetSyncStatus(status.AddressId)
				if err != nil {
					log.Println("ERROR SYNC:", err)
				}
				return
			}

			status.Cursor = cursorTxs.NextCursor
		}
	}
}
