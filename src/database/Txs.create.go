package database

import (
	"context"
)

func (d *Database) TxCreate(hash string, data string) (*Tx, error) {
	d.conn.Exec(
		context.Background(),
		`INSERT INTO "Txs" (hash, data) VALUES($1,$2)`,
		hash, data,
	)

	return d.TxsGetByHash(hash)
}
