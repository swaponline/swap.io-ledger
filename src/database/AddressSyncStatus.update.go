package database

import (
	"context"
)

func (d *Database) AddressSyncStatusUpdateCursor(
	addressId int,
	cursor string,
) error {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(
		context.Background(),
		`update  "Address_sync_status"
		 set cursor_id = $2
		 where address_id = $1
		`,
		addressId, cursor,
	)

	return err
}
func (d *Database) AddressSyncStatusUpdateSyncStatus(
	addressId int,
	syncStatus int,
) error {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(
		context.Background(),
		`update  "Address_sync_status"
		 set sync = $2
		 where address_id = $1
		`,
		addressId, syncStatus,
	)

	return err
}
