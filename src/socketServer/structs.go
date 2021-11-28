package socketServer

import "swap.io-ledger/src/database"

type TxNotification struct {
	Tx *database.Tx
	UsersIds []int
}
