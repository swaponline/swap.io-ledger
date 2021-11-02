package TxsManager

import "swap.io-ledger/src/database"

func (tm *TxsManager) CreateTx(
	hash string,
	data string,
) (*database.Tx, error) {
	return tm.database.TxCreate(hash, data)
}
