package database

import (
	"context"
	"log"
)

func (d *Database) UsersCreate (
	pubKey string,
)  {
	_, err := d.conn.Exec(
		context.Background(),
		"INSERT INTO Users (pub_key) values($1)",
		pubKey,
	)
	log.Println(err)
}