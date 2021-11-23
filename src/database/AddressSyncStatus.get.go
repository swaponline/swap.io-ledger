package database

import (
	"context"
)

func (d *Database) AddressSyncStatusGetById(addressId int) (*AddressSyncStatus, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}

	addressSyncStatus := new(AddressSyncStatus)
	err = conn.QueryRow(
		context.Background(),
		`select 
			address_id,
			sync,
			cursor_id,
			address,
			n.name
		from "Address_sync_status" ass 
		inner join "Users_addresses" ua
			on ass.address_id = ua.id and ass.address_id = $1
		inner join "Coins" c
			on ua.coin_id = c.id
		inner join "Networks" n
			on c.network_id = n.id
		`, addressId,
	).Scan(
		&addressSyncStatus.AddressId,
		&addressSyncStatus.Sync,
		&addressSyncStatus.Cursor,
		&addressSyncStatus.Address,
		&addressSyncStatus.Network,
	)
	if err != nil {
		return nil, err
	}

	return addressSyncStatus, nil
}
func (d *Database) AddressSyncStatusGetNotSyncAddresses() ([]AddressSyncStatus, error) {
	statusesRows, err := d.pool.Query(
		context.Background(),
		`select 
			address_id,
			sync,
			cursor_id,
			address,
			n.name
		from "Address_sync_status" ass 
		inner join "Users_addresses" ua
			on sync = 0 and ass.address_id = ua.id
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
		statuses = append(statuses, status)
	}

	return statuses, err
}
