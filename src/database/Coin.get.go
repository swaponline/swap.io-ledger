package database

import (
	"context"
)

func (d *Database) CoinGetByName(
	name string,
) (*Coin, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	coin := new(Coin)
	err = conn.QueryRow(
		context.Background(),
		`select id, name from "Coins" where name = $1`,
		name,
	).Scan(
		&coin.Id,
		&coin.Name,
	)

	return coin, err
}
