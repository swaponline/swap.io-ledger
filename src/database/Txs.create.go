package database

import (
	"context"
)

func (d *Database) TxCreate(hash string, data string) (*Tx, error) {
	_, err := d.conn.Exec(
		context.Background(),
		`INSERT INTO Transactions (hash, data) VALUES($1,$2)`,
		hash, data,
	)
	if err != nil {
		return nil, err
	}

	return d.TxsGetByHash(hash)
}
