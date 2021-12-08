package database

import "context"

func (d *Database) UserSpendsGetSummaryUserSpends(userId int) (*SummaryUserSpends, error) {
	conn, err := d.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}

	summaryUserSpends := SummaryUserSpends{}
	// todo: add sql
	err = conn.QueryRow(context.Background(), ``, userId).Scan(
		&summaryUserSpends.Network,
		&summaryUserSpends.Coin,
		&summaryUserSpends.value,
	)
	if err != nil {
		return nil, err
	}

	return &summaryUserSpends, err
}
