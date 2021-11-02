package CoinsManager

import "swap.io-ledger/src/database"

func (cm *CoinsManager) getCoin(name string) (*database.Coin, error) {
	return cm.database.CoinGetByName(name)
}
