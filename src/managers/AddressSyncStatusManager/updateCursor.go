package AddressSyncStatusManager

func (assm *AddressSyncStatusManager) updateCursor(
	addressId int,
	cursor string,
) error {
	return assm.database.AddressSyncStatusUpdateCursor(
		addressId,
		cursor,
	)
}
