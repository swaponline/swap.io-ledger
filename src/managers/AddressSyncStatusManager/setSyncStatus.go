package AddressSyncStatusManager

func (assm *AddressSyncStatusManager) setSyncStatus(
	addressId int,
) error {
	return assm.database.AddressSyncStatusUpdateSyncStatus(
		addressId,
		1,
	)
}
