package TxsManager

func (tm *TxsManager) GetTxsByUserId(userId int) (map[string][]string, error) {
	return tm.database.GetTxsByUserId(userId)
}
