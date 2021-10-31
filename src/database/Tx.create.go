package database

import (
	"context"
	"encoding/json"
	"swap.io-ledger/src/agentHandler"
)

func (d *Database) TxCreate(tx *agentHandler.Transaction) error {
	if saveData, err := json.Marshal(tx); err == nil {
		_, err := d.conn.Exec(
			context.Background(),
			`INSERT INTO Transactions (hash, data) VALUES($1,$2)`,
			tx.Hash, string(saveData),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
