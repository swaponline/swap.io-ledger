package database

import "context"

func (d *Database) TxsGetByHash(hash string) (*Tx, error) {
	tx := new(Tx)
	err := d.conn.QueryRow(
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