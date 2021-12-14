package database

import (
	"context"
)

func (d *Database) UsersSpendsCreate(
	txId int,
	txSpendIndex int,
	userAddressId int,
	value string,
) error {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(
		context.Background(),
		`insert into "Users_spends" 
			(tx_id, tx_spend_index, user_address_id, value)
			values ($1, $2, $3, $4)
		`,
		txId, txSpendIndex, userAddressId, value,
	)

	return err
}
