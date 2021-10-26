package txsHandler

import (
    "log"
    "time"

    "swap.io-ledger/src/agentHandler"
    "swap.io-ledger/src/database"
    "swap.io-ledger/src/usersManager"
)

type TxsHandler struct {
    network string
    txSource    chan *agentHandler.Transaction
    txIsReceive chan struct{}
    realtimeTxs chan *agentHandler.Transaction
    database *database.Database
}
type Config struct {
    Network string
    TxSource chan *agentHandler.Transaction
    TxIsReceive chan struct{}
    Database *database.Database
    UsersManager *usersManager.UsersManager
}

func InitialiseTxsHandler(config Config) *TxsHandler {
    return &TxsHandler{
        network: config.Network,
        txSource: config.TxSource,
        txIsReceive: config.TxIsReceive,
        database: config.Database,
        realtimeTxs: make(chan *agentHandler.Transaction),
    }
}

func (t *TxsHandler) Start() {
    for {
        tx := <-t.txSource
        for err := t.database.SaveTx(tx); err != nil; {
            <-time.After(time.Second)
        }
        t.txIsReceive <- struct{}{}
        log.Println("receive", tx.Hash)
        //t.realtimeTxs <- tx
    }
}
func (*TxsHandler) Status() error {
    return nil
}
func (*TxsHandler) Stop() error {
    return nil
}
