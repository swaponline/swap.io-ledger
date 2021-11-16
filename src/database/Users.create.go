package database

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func (d *Database) UsersCreate (
	pubKey string,
) (int,error) {
	var newUserId int
	err := d.pool.QueryRow(
		context.Background(),
		`insert into "Users" (pub_key) values($1)`,
		pubKey,
	).Scan(&newUserId)

	return newUserId, err
}
func (d *Database) UsersCreateByPubKeyAndAddresses(
	pubKey string,
	addresses []CreateUserAddressData,
) (int, []int, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return 0, nil, err
	}
	dTx, err := conn.BeginTx(context.Background(), pgx.TxOptions{});
	if err != nil {
		return 0, nil, err
	}
	defer conn.Release()
	defer dTx.Rollback(context.Background());

	_, err = dTx.Exec(
		context.Background(),
		`insert into "Users" (pub_key) values($1)`,
		pubKey,
	)
	if err != nil {
		return 0, nil, err
	}

	var newUserId int
	newAddressesIds := make([]int, 0)
	row := dTx.QueryRow(
		context.Background(),
		`select id from "Users" where pub_key = $1`,
		pubKey,
	)
	err = row.Scan(&newUserId)
	if err != nil {
		return 0, nil, err
	}

	for _, address := range addresses {
		// todo: errors handle
		var coinId int
		row = dTx.QueryRow(
			context.Background(),
			`select c.id from "Networks" n inner join "Coins" c
			 on n.id = c.network_id and 
                n.name = $1 and 
                c.name = $2`,
			 address.Network, address.Coin,
		)
		err := row.Scan(&coinId)
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
			coinId,newUserId,address,
		).Scan(&newAddressId)
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
			newAddressId,0,"null",
		)
		if err != nil {
			return 0, nil, err
		}

		newAddressesIds = append(
			newAddressesIds,
			newAddressId,
		)
	}

	err = dTx.Commit(context.Background())
	if err != nil {
		return 0, nil, err
	}

	return newUserId, newAddressesIds, nil
}