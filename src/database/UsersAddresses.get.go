package database

import (
	"context"
)

func (d *Database) UsersAddressesGetByCoinIdAndAddress(
	coinId int,
	address string,
) (*UserAddress, error) {
	userAddress := new(UserAddress)
	err := d.pool.QueryRow(
		context.Background(),
		`select 
			id,
			coin_id,
			user_id,
			address
		from "Users_addresses" where coin_id = $1 and address = $2`,
		coinId, address,
	).Scan(
		&userAddress.Id,
		&userAddress.CoinId,
		&userAddress.UserId,
		&userAddress.Address,
	)

	return userAddress, err
}