package database

import "context"

func (d *Database) AddressSyncStatusGetNotSync() ([]AddressSyncStatus, error) {
	statusesRows, err := d.conn.Query(
		context.Background(),
		`select address_id, sync, cursor from "Address_sync_status" where sync = 0`,
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
		)
		if err != nil {
			return nil, err
		}
	}

	return statuses, err
}
