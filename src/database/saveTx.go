package database

import (
	"context"

	"swap.io-ledger/src/agentHandler"
)

func (d *Database) SaveTx(tx agentHandler.Transaction) error {
    d.conn.Exec(
        context.Background(),
        ``,
    )

    return nil
}
