package database

import "context"

func (d *Database) TxsGetByHash(hash string) (*Tx, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	tx := new(Tx)
	err = conn.QueryRow(
		context.Background(),
		`select id, hash, data from "Txs" where hash = $1`,
		hash,
	).Scan(
		&tx.Id,
		&tx.Hash,
		&tx.Data,
	)

	return tx, err
}