package database

import (
	"context"
	"swap.io-ledger/src/managers/UsersAdressesManager"
)

func (d *Database) UsersAddressesGetByAddress(
	address string,
) (*UsersAdressesManager.UserAddress, error) {
	rowData := d.conn.QueryRow(
		context.Background(),
		`select 
			id,
			coin_id,
			user_id,
			address
		from "Users_addresses" where address = $1`,
		address,
	)
	userAddress := new(UsersAdressesManager.UserAddress)
	err := rowData.Scan(
		&userAddress.Id,
		&userAddress.CoinId,
		&userAddress.UserId,
		&userAddress.Address,
	)
	if err != nil {
		return nil, err
	}

	return nil, err
}