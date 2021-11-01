package txsHandler

import (
    "log"
    "swap.io-ledger/src/managers/UsersManager"
    "time"

    "swap.io-ledger/src/agentHandler"
    "swap.io-ledger/src/database"
)

type TxsHandler struct {
    network string
    txSource    chan *agentHandler.AgentTx
    txIsReceive chan struct{}
    realtimeTxs chan *agentHandler.AgentTx
    database *database.Database
}
type Config struct {
    Network string
    TxSource chan *agentHandler.AgentTx
    TxIsReceive chan struct{}
    Database *database.Database
    UsersManager *UsersManager.UsersManager
}

func InitialiseTxsHandler(config Config) *TxsHandler {
    return &TxsHandler{
        network: config.Network,
        txSource: config.TxSource,
        txIsReceive: config.TxIsReceive,
        database: config.Database,
        realtimeTxs: make(chan *agentHandler.AgentTx),
    }
}

func (t *TxsHandler) Start() {
    for {
        tx := <-t.txSource
        for err := t.database.TxCreate(tx); err != nil; {
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
