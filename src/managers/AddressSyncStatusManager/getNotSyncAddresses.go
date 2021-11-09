package AddressSyncStatusManager

import "swap.io-ledger/src/database"

func (assm *AddressSyncStatusManager) getNotSyncAddresses() (
	[]database.AddressSyncStatus,
	error,
) {
	return assm.database.AddressSyncStatusGetNotSyncAddresses()
}