package UsersSpendsManager

import "swap.io-ledger/src/database"

func (usm *UsersSpendsManager) GetUserBalances(userId int) (
	[]database.UserBalance,
	error,
) {
	return usm.database.UserSpendsGetUserBalances(userId)
}
