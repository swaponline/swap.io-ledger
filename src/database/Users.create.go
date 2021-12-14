package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func (d *Database) UsersCreate(
	pubKey string,
) (int, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var newUserId int
	err = conn.QueryRow(
		context.Background(),
		`insert into "Users" (pub_key) values($1)`,
		pubKey,
	).Scan(&newUserId)

	return newUserId, err
}
func (d *Database) UsersCreateByPubKeyAndAddresses(
	pubKey string,
	addresses []CreateUserAddressData,
	beforeCommit func() error,
) (int, []int, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return 0, nil, err
	}
	defer conn.Release()

	dTx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return 0, nil, err
	}
	defer dTx.Rollback(context.Background())

	var newUserId int
	err = dTx.QueryRow(
		context.Background(),
		`insert into "Users" (pub_key) values($1) returning id`,
		pubKey,
	).Scan(&newUserId)
	if err != nil {
		return 0, nil, err
	}
	log.Println("tx user created", newUserId)

	newAddressesIds := make([]int, 0)
	for _, address := range addresses {
		// todo: errors handle
		var coinId int
		err := dTx.QueryRow(
			context.Background(),
			`select c.id from "Networks" n inner join "Coins" c
			 on n.id = c.network_id and 
                n.name = $1 and 
                c.name = $2`,
			address.Network, address.Coin,
		).Scan(&coinId)
		log.Println("tx find coinId", coinId, err)
		if err != nil {
			return 0, nil, err
		}

		var newAddressId int
		err = dTx.QueryRow(
			context.Background(),
			`insert into "Users_addresses" (
				coin_id,
				user_id,
				address
			) values($1,$2,$3) returning id`,
			coinId, newUserId, address.Address,
		).Scan(&newAddressId)
		log.Println("tx find newAddressId", newAddressId)
		if err != nil {
			return 0, nil, err
		}

		_, err = dTx.Exec(
			context.Background(),
			`insert into "Address_sync_status" (
				address_id,
				sync,
				cursor_id
			) values($1,$2,$3)`,
			newAddressId, 0, "null",
		)
		if err != nil {
			return 0, nil, err
		}

		newAddressesIds = append(
			newAddressesIds,
			newAddressId,
		)
	}

	err = beforeCommit()
	if err != nil {
		return 0, nil, err
	}

	err = dTx.Commit(context.Background())
	if err != nil {
		return 0, nil, err
	}

	return newUserId, newAddressesIds, nil
}
