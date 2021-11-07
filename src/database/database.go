package database

import (
	"context"
	"github.com/jackc/pgtype"
	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/jackc/pgx/v4"
	"log"
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
	conn.ConnInfo().RegisterDataType(pgtype.DataType{
		Value: &shopspring.Numeric{},
		Name:  "numeric",
		OID:   pgtype.NumericOID,
	})

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
