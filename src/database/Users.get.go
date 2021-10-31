package database

import (
	"context"
	"swap.io-ledger/src/managers/UsersManager"
)

func (d *Database) UsersGetById(id int) (*UsersManager.User, error) {
	rowData := d.conn.QueryRow(
		context.Background(),
		`select id, pub_key from "Users" where id = $1`,
		id,
	)
	user := new(UsersManager.User)
	err := rowData.Scan(
		&user.Id,
		&user.PubKey,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}