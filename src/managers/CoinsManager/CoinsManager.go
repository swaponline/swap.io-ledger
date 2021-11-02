package CoinsManager

import "swap.io-ledger/src/database"

type CoinsManager struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseCoinsManager(config Config) *CoinsManager {
	return &CoinsManager{
		database: config.Database,
	}
}

func (*Config) Start() {}
func (*Config) Status() error {
	return nil
}
func (*Config) Stop() error {
	return nil
}
