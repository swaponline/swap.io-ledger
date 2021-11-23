package AddressSyncStatusManager

import "swap.io-ledger/src/database"

func (as *AddressSyncStatusManager) CreateAddressSyncStatus(
	addressId int,
	sync int,
	cursorId string,
) (*database.AddressSyncStatus, error) {
	return as.database.AddressSyncStatusCreate(addressId, sync, cursorId)
}