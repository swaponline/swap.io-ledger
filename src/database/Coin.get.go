package database

import (
	"context"
)

func (d *Database) CoinGetByName(
	name string,
) (*Coin, error) {
	coin := new(Coin)
	err := d.pool.QueryRow(
		context.Background(),
		`select id, name from "Coins" where name = $1`,
		name,
	).Scan(
		&coin.Id,
		&coin.Name,
	)

	return coin, err
}
