package AddressSyncStatusManager

import "swap.io-ledger/src/database"

func (assm *AddressSyncStatusManager) GetNotSyncAddresses() (
	[]database.AddressSyncStatus,
	error,
) {
	return assm.database.AddressSyncStatusGetNotSyncAddresses()
}