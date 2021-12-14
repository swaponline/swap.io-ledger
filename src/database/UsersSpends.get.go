package database

import (
	"context"
)

func (d *Database) UserSpendsGetUserBalances(userId int) (
	[]UserBalance,
	error,
) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `
		select n.name, c.name, sum(us.value) from "Users" u 
			inner join "Users_addresses" ua
				on u.id = $1 and 
 				   u.id = ua.user_id
			inner join "Coins" c 
				on ua.coin_id = c.id
			inner join "Networks" n
				on c.network_id = n.id
			inner join "Users_spends" us
				on ua.id = us.user_address_id
		group by n.name, c.name
	`, userId)
	if err != nil {
		return nil, err
	}

	allUserBalances := make([]UserBalance, 0)
	for rows.Next() {
		userBalance := UserBalance{}
		err := rows.Scan(
			&userBalance.Network,
			&userBalance.Coin,
			&userBalance.Balance,
		)
		if err != nil {
			return nil, err
		}

		allUserBalances = append(
			allUserBalances,
			userBalance,
		)
	}

	return allUserBalances, err
}
