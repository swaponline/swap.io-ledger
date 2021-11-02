package database

import (
	"context"
)

func (d *Database) UsersAddressesGetByAddress(
	address string,
) (*UserAddress, error) {
	userAddress := new(UserAddress)
	err := d.conn.QueryRow(
		context.Background(),
		`select 
			id,
			coin_id,
			user_id,
			address
		from "Users_addresses" where address = $1`,
		address,
	).Scan(
		&userAddress.Id,
		&userAddress.CoinId,
		&userAddress.UserId,
		&userAddress.Address,
	)

	return userAddress, err
}