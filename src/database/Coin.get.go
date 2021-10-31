package database

import (
	"context"
	"swap.io-ledger/src/managers/CoinsManager"
)

func (d *Database) CoinGetByName(
	name string,
) (*CoinsManager.Coin, error) {
	rowData := d.conn.QueryRow(
		context.Background(),
		`select id, name from "Coin" where name = $1`,
		name,
	)
	coin := new(CoinsManager.Coin)
	err := rowData.Scan(
		&coin.Id,
		&coin.Name,
	)
	if err != nil {
		return nil, err
	}

	return coin, err
}
