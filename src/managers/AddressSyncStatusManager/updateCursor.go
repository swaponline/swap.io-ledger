package AddressSyncStatusManager

func (assm *AddressSyncStatusManager) UpdateCursor(
	addressId int,
	cursor string,
) error {
	return assm.database.AddressSyncStatusUpdateCursor(
		addressId,
		cursor,
	)
}
