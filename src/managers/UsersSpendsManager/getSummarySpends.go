package UsersSpendsManager

import "swap.io-ledger/src/database"

func (usm *UsersSpendsManager) GetSummarySpends(userId int) (
	*database.SummaryUserSpends,
	error,
) {
	return usm.database.UserSpendsGetSummaryUserSpends(userId)
}
