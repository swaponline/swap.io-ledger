package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"swap.io-ledger/src/config"
)

type Database struct {
    conn *pgx.Conn
}

func InitialiseDatabase() *Database {
    conn, err := pgx.Connect(context.Background(), config.POSTGRESS_URL)
    if err != nil {
        log.Panicln(err)
    }

    return &Database{
        conn: conn,
    }
}

func (*Database) Start() {}
func (*Database) Status() error {
    return nil
}
func (d *Database) Stop() error {
    return d.conn.Close(context.Background())
}
