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