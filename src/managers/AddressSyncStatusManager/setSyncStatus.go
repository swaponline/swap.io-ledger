package AddressSyncStatusManager

func (assm *AddressSyncStatusManager) SetSyncStatus(
	addressId int,
) error {
	return assm.database.AddressSyncStatusUpdateSyncStatus(
		addressId,
		1,
	)
}
