package AddressSyncStatusManager

import "swap.io-ledger/src/database"

func (ass *AddressSyncStatusManager) GetByAddressId(addressId int) (
	*database.AddressSyncStatus,
	error,
) {
	return ass.database.AddressSyncStatusGetById(addressId)
}
