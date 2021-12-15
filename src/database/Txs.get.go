package database

import "context"

func (d *Database) TxsGetByHash(hash string) (*Tx, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	tx := new(Tx)
	err = conn.QueryRow(
		context.Background(),
		`select id, hash, data from "Txs" where hash = $1`,
		hash,
	).Scan(
		&tx.Id,
		&tx.Hash,
		&tx.Data,
	)

	return tx, err
}
func (d *Database) GetTxsByUserId(userId int) (map[string][]string, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rowTxs, err := conn.Query(context.Background(), `
		select n.name, t.data from (
			select DISTINCT n.id n_id, t.id t_id from "Users_addresses" ua 
				inner join "Users_spends" us 
					on ua.user_id = $1 and
					   ua.id = us.user_address_id
				inner join "Txs" t
					on us.tx_id = t.id
				inner join "Coins" c
					on ua.coin_id = c.id
				inner join "Networks" n
					on c.network_id = n.id
		) compactTxInfo 
			inner join "Networks" n
				on compactTxInfo.n_id = n.id
			inner join "Txs" t
				on compactTxInfo.t_id = t.id
	`, userId)
	if err != nil {
		return nil, err
	}

	userTxs := make(map[string][]string)
	for rowTxs.Next() {
		var network string
		var data string
		err := rowTxs.Scan(&network, &data)
		if err != nil {
			return nil, err
		}

		userTxs[network] = append(userTxs[network], data)
	}
	if rowTxs.Err() != nil {
		return nil, rowTxs.Err()
	}

	return userTxs, nil
}
