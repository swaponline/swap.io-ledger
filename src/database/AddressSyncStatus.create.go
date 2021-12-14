package database

import "context"

func (d *Database) AddressSyncStatusCreate(
	addressId int,
	sync int,
	cursor string,
) (*AddressSyncStatus, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	_, err = conn.Exec(
		context.Background(),
		`insert into "Address_sync_status" (address_id, sync, cursor)
			values($1,$2,$3)
		`,
		addressId, sync, cursor,
	)
	if err != nil {
		return nil, err
	}

	status := AddressSyncStatus{
		AddressId: addressId,
		Sync:      sync,
		Cursor:    cursor,
	}

	return &status, err
}
