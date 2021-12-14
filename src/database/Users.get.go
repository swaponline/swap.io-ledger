package database

import (
	"context"
)

func (d *Database) UsersGetById(id int) (*User, error) {
	user := new(User)
	err := d.pool.QueryRow(
		context.Background(),
		`select id, pub_key from "Users" where id = $1`,
		id,
	).Scan(
		&user.Id,
		&user.PubKey,
	)

	return user, err
}
func (d *Database) UsersGetByPubKey(pubKey string) (*User, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	user := new(User)
	err = conn.QueryRow(
		context.Background(),
		`select id, pub_key from "Users" where pub_key = $1`,
		pubKey,
	).Scan(&user.Id, &user.PubKey)
	if err != nil {
		return nil, err
	}

	return user, nil
}
