package agentHandler

import "swap.io-ledger/src/txsHandler"

type CursorTxs struct {
	Cursor       string
	NextCursor   string
	Transactions []*txsHandler.NonHandledTx
}
