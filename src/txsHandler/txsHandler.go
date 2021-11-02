package txsHandler

import (
    "encoding/json"
    "log"
    "swap.io-ledger/src/agentHandler"
    "swap.io-ledger/src/managers/TxsManager"
    "swap.io-ledger/src/managers/UsersAdressesManager"
    "swap.io-ledger/src/managers/UsersSpendsManager"
)

type TxsHandler struct {
    network string
    aTxSource    chan *agentHandler.AgentTx
    aTxIsReceive chan struct{}
    realtimeTxs chan *agentHandler.AgentTx
    txsManager            *TxsManager.TxsManager
    usersAddressesManager *UsersAdressesManager.UsersAddressesManager
    usersSpendsManager    *UsersSpendsManager.UsersSpendsManager
}
type Config struct {
    Network string
    aTxSource chan *agentHandler.AgentTx
    aTxIsReceive chan struct{}
    TxsManager            *TxsManager.TxsManager
    UsersAddressesManager *UsersAdressesManager.UsersAddressesManager
    UsersSpendsManager    *UsersSpendsManager.UsersSpendsManager
}

func InitialiseTxsHandler(config Config) *TxsHandler {
    return &TxsHandler{
        network: config.Network,
        aTxSource: config.aTxSource,
        aTxIsReceive: config.aTxIsReceive,
        realtimeTxs: make(chan *agentHandler.AgentTx),
    }
}

func (t *TxsHandler) Start() {
    // todo: handle all errors
    for {
        aTx := <-t.aTxSource

        aTxData, err := json.Marshal(aTx)
        if err != nil {
            t.aTxIsReceive <- struct{}{}
            continue
        }

        tx, err := t.txsManager.CreateTx(
            aTx.Hash,
            string(aTxData),
        )
        if err != nil {
            t.aTxIsReceive <- struct{}{}
            continue
        }

        //aTx.Journal

        t.aTxIsReceive <- struct{}{}
        log.Println("receive", aTx.Hash)
        //t.realtimeTxs <- tx
    }
}
func (*TxsHandler) Status() error {
    return nil
}
func (*TxsHandler) Stop() error {
    return nil
}
