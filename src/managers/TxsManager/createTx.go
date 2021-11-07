package TxsManager

import (
	"swap.io-ledger/src/database"
	"time"
)

func (tm *TxsManager) CreateTx(
	hash string,
	data string,
) *database.Tx {
	for {
		tx, err := tm.database.TxCreate(hash, data)
		if err != nil {
			time.After(time.Second)
			continue
		}

		return tx
	}
}
