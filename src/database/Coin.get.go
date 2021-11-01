package database

import (
	"context"
)

func (d *Database) CoinGetByName(
	name string,
) (*Coin, error) {
	rowData := d.conn.QueryRow(
		context.Background(),
		`select id, name from "Coin" where name = $1`,
		name,
	)
	coin := new(Coin)
	err := rowData.Scan(
		&coin.Id,
		&coin.Name,
	)
	if err != nil {
		return nil, err
	}

	return coin, err
}
