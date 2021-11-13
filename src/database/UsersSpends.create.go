package database

import (
	"context"
)

func (d *Database) UsersSpendsCreate(
	txId int,
	txSpendIndex int,
	userAddressId int,
	value string,
) error  {
	_, err := d.pool.Exec(
		context.Background(),
		`insert into "Users_spends" 
			(tx_id, tx_spend_index, user_address_id, value)
			values ($1, $2, $3, $4)
		`,
		txId, txSpendIndex, userAddressId, value,
	)

	return err
}