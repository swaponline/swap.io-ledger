package database

import "context"

func (d *Database) AddressSyncStatusGetNotSyncAddresses() ([]AddressSyncStatus, error) {
	statusesRows, err := d.conn.Query(
		context.Background(),
		`select 
			address_id,
			sync,
			cursor,
			address,
			n.name
		from "Address_sync_status" ass 
		inner join "Users_addresses" ua
			on sync = 0 and ass.address_id = us.id
		inner join "Coins" c
			on ua.coin_id = c.id
		inner join "Networks" n
			on c.network_id = n.id
		`,
	)
	if err != nil {
		return nil, err
	}

	statuses := make([]AddressSyncStatus, 0)
	for statusesRows.Next() {
		status := AddressSyncStatus{}
		err := statusesRows.Scan(
			&status.AddressId,
			&status.Sync,
			&status.Cursor,
			&status.Address,
			&status.Network,
		)
		if err != nil {
			return nil, err
		}
	}

	return statuses, err
}
