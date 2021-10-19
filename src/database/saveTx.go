package database

import (
	"context"
	"encoding/json"
	"log"

	"swap.io-ledger/src/agentHandler"
)

func (d *Database) SaveTx(tx *agentHandler.Transaction) error {
    if saveData, err := json.Marshal(tx); err != nil {
		log.Println("write new tx", string(saveData))
        d.conn.Exec(
            context.Background(),
            `INSERT INTO transactions (hash, data) values('','')`,
            tx.Hash, string(saveData),
        )
    }

    return nil
}
