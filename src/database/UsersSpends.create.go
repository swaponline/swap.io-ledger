package database

import "context"

func (d *Database) UsersSpendsCreate(
	txId int,
	txWiringIndex int,
	userAddressId int,
	value int,
) error  {
	_, err := d.conn.Exec(
		context.Background(),
		`insert into "Users_spends" 
			(tx_id, tx_wiring_index, user_address_id, value)
			values ($1, $2, $3, $4)
		`,
		txId, txWiringIndex, userAddressId, value,
	)

	return err
}