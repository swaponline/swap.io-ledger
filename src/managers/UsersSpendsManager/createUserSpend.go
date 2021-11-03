package UsersSpendsManager

func (usm *UsersSpendsManager) CreateUserSpend(
	data CreateUserSpendData,
) error {
	return usm.database.UsersSpendsCreate(
		data.TxId,
		data.TxSpendIndex,
		data.UsersAddressId,
		data.Value,
	)
}