package AgentHandler

import (
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/txsHandler"
)

type CursorTxs struct {
	Cursor       string
	NextCursor   string
	Transactions []*txsHandler.NonHandledTx
}
type TxNotification struct {
	Tx *database.Tx
	UsersIds []int
}
