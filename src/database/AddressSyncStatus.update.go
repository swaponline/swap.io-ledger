package database

import "context"

func (d *Database) AddressSyncStatusUpdateCursor(
	addressId int,
	cursor string,
) error {
	_, err := d.conn.Exec(
		context.Background(),
		`update  "Address_sync_status"
		 set cursor = $2
		 where id = $1
		`,
		addressId, cursor,
	)

	return err
}
func (d *Database) AddressSyncStatusUpdateSyncStatus(
	addressId int,
	syncStatus int,
) error {
	_, err := d.conn.Exec(
		context.Background(),
		`update  "Address_sync_status"
		 set sync = $2
		 where id = $1
		`,
		addressId, syncStatus,
	)

	return err
}