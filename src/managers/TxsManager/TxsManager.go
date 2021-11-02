package TxsManager

import "swap.io-ledger/src/database"

type TxsManager struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseTxsManager(config Config) *TxsManager {
	return &TxsManager{
		database: config.Database,
	}
}

func (*TxsManager) Start() {}
func (*TxsManager) Status() error {
	return nil
}
func (*TxsManager) Stop() error {
	return nil
}
