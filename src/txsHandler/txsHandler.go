package txsHandler

import (
    "encoding/json"
    "log"
    "swap.io-ledger/src/agentHandler"
    "swap.io-ledger/src/managers/CoinsManager"
    "swap.io-ledger/src/managers/TxsManager"
    "swap.io-ledger/src/managers/UsersAdressesManager"
    "swap.io-ledger/src/managers/UsersSpendsManager"
)

type TxsHandler struct {
    network string
    aTxSource    chan *agentHandler.AgentTx
    aTxIsReceive chan struct{}
    realtimeTxs  chan *agentHandler.AgentTx
    txsManager            *TxsManager.TxsManager
    coinsManager          *CoinsManager.CoinsManager
    usersAddressesManager *UsersAdressesManager.UsersAddressesManager
    usersSpendsManager    *UsersSpendsManager.UsersSpendsManager
}
type Config struct {
    Network string
    ATxSource chan *agentHandler.AgentTx
    ATxIsReceive chan struct{}
    TxsManager            *TxsManager.TxsManager
    CoinsManager          *CoinsManager.CoinsManager
    UsersAddressesManager *UsersAdressesManager.UsersAddressesManager
    UsersSpendsManager    *UsersSpendsManager.UsersSpendsManager
}

func InitialiseTxsHandler(config Config) *TxsHandler {
    return &TxsHandler{
        network: config.Network,
        aTxSource: config.ATxSource,
        aTxIsReceive: config.ATxIsReceive,
        realtimeTxs: make(chan *agentHandler.AgentTx),
    }
}

func (th *TxsHandler) Start() {
    // todo: handle all errors
    for {
        aTx := <-th.aTxSource

        aTxData, err := json.Marshal(aTx)
        if err != nil {
            th.aTxIsReceive <- struct{}{}
            continue
        }

        tx, err := th.txsManager.CreateTx(
            aTx.Hash,
            string(aTxData),
        )
        if err != nil {
            th.aTxIsReceive <- struct{}{}
            continue
        }

        for _, spendsInfo := range aTx.Journal {
            coin, err := th.coinsManager.GetCoin(
                spendsInfo.Asset.Id,
            )
            if err != nil {
                continue
            }
            for spendIndex, spend := range spendsInfo.Entries {
                userAddress, err := th.usersAddressesManager.GetUserAddress(
                    coin.Id, spend.Wallet,
                )
                if err != nil {
                    continue
                }
                th.usersSpendsManager.CreateUserSpend(
                    UsersSpendsManager.CreateUserSpendData{
                        TxId: tx.Id,
                        TxSpendIndex: spendIndex,
                        UsersAddressId: userAddress.Id,
                        Value: 100,
                    },
                )
            }
        }

        //aTx.Journal

        th.aTxIsReceive <- struct{}{}
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
