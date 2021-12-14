package database

import (
	"context"
)

func (d *Database) UsersAddressesGetByCoinIdAndAddress(
	coinId int,
	address string,
) (*UserAddress, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	userAddress := new(UserAddress)
	err = conn.QueryRow(
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
