package database

import (
	"context"
)

func (d *Database) TxCreate(hash string, data string) (*Tx, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	conn.Exec(
		context.Background(),
		`INSERT INTO "Txs" (hash, data) VALUES($1,$2)`,
		hash, data,
	)

	return d.TxsGetByHash(hash)
}
