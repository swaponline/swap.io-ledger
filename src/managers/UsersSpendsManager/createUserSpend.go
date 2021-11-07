package UsersSpendsManager

func (usm *UsersSpendsManager) CreateUserSpend(
	data CreateUserSpendData,
) error {
	// todo: error handle
	return usm.database.UsersSpendsCreate(
		data.TxId,
		data.TxSpendIndex,
		data.UsersAddressId,
		data.Value,
	)
}