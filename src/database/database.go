package database

import (
	"context"
	"github.com/jackc/pgtype"
	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"swap.io-ledger/src/config"
)

type Database struct {
    pool *pgxpool.Pool
}

func InitialiseDatabase() *Database {
	pgxPoolConfig, err := pgxpool.ParseConfig(config.POSTGRESS_URL)
	if err != nil {
		log.Panicln(err)
	}
	pgxPoolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &shopspring.Numeric{},
			Name:  "numeric",
			OID:   pgtype.NumericOID,
		})
		return nil
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), pgxPoolConfig);

    return &Database{
        pool: pool,
    }
}

func (*Database) Start() {}
func (*Database) Status() error {
    return nil
}
func (d *Database) Stop() error {
    d.pool.Close();
	return nil;
}
