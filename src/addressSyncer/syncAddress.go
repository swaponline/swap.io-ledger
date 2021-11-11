package addressSyncer

import (
	"log"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/database"
	"time"
)

func (as *AddressSyncer) SyncAddress(
	status *database.AddressSyncStatus,
) {
	if agentHandlerInstance, ok := as.agentHandlers[status.Network]; ok {
		for {
			var cursorTxs *agentHandler.CursorTxs
			var err error
			if status.Cursor == "null" {
				cursorTxs, err = agentHandlerInstance.GetFirstTxsCursor(status.Address)
			} else {
				cursorTxs, err = agentHandlerInstance.GetTxsCursor(status.Cursor)
			}
			if err != nil {
				log.Println("ERROR:", err)
				<-time.After(time.Second)
				continue
			}

			for _, aTx := range cursorTxs.Transactions {
				as.txsHandler.TxHandle(aTx)
			}

			if cursorTxs.NextCursor != "null" {
				err := as.addressSyncStatusManager.UpdateCursor(
					status.AddressId,
					cursorTxs.NextCursor,
				)
				if err != nil {
					log.Println("ERROR:", err)
				}
			} else {
				err := as.addressSyncStatusManager.SetSyncStatus(status.AddressId)
				if err != nil {
					log.Println("ERROR:", err)
				}
				return
			}

			cursorTxs.Cursor = cursorTxs.NextCursor
		}
	}
}
