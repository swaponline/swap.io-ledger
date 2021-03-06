package txsHandler

import (
	"encoding/json"
	"log"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/managers/CoinsManager"
	"swap.io-ledger/src/managers/TxsManager"
	"swap.io-ledger/src/managers/UsersAdressesManager"
	"swap.io-ledger/src/managers/UsersSpendsManager"
	"swap.io-ledger/src/serviceRegistry"
)

type TxsHandler struct {
	txsManager            *TxsManager.TxsManager
	coinsManager          *CoinsManager.CoinsManager
	usersAddressesManager *UsersAdressesManager.UsersAddressesManager
	usersSpendsManager    *UsersSpendsManager.UsersSpendsManager
}
type Config struct {
	TxsManager            *TxsManager.TxsManager
	CoinsManager          *CoinsManager.CoinsManager
	UsersAddressesManager *UsersAdressesManager.UsersAddressesManager
	UsersSpendsManager    *UsersSpendsManager.UsersSpendsManager
}

func InitialiseTxsHandler(config Config) *TxsHandler {
	return &TxsHandler{
		txsManager:            config.TxsManager,
		coinsManager:          config.CoinsManager,
		usersAddressesManager: config.UsersAddressesManager,
		usersSpendsManager:    config.UsersSpendsManager,
	}
}
func Register(reg *serviceRegistry.ServiceRegistry) {
	var txsManager *TxsManager.TxsManager
	err := reg.FetchService(&txsManager)
	if err != nil {
		log.Panicln(err)
	}

	var coinsManager *CoinsManager.CoinsManager
	err = reg.FetchService(&coinsManager)
	if err != nil {
		log.Panicln(err)
	}

	var usersAddressesManager *UsersAdressesManager.UsersAddressesManager
	err = reg.FetchService(&usersAddressesManager)
	if err != nil {
		log.Panicln(err)
	}

	var usersSpendsManager *UsersSpendsManager.UsersSpendsManager
	err = reg.FetchService(&usersSpendsManager)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(
		InitialiseTxsHandler(Config{
			TxsManager:            txsManager,
			CoinsManager:          coinsManager,
			UsersAddressesManager: usersAddressesManager,
			UsersSpendsManager:    usersSpendsManager,
		}),
	)
}

// TxHandle added spends registered users to db and return registered users
func (th *TxsHandler) TxHandle(
	nonTx *NonHandledTx,
) (*database.Tx, []int) {
	// todo: handle all errors
	nonTxData, _ := json.Marshal(nonTx)

	//log.Println(string(aTxData), th)
	tx := th.txsManager.CreateTx(
		nonTx.Hash,
		string(nonTxData),
	)
	participateUserIdsMap := make(map[int]struct{})

	for _, spendsInfo := range nonTx.Journal {
		coin, err := th.coinsManager.GetCoin(
			spendsInfo.Asset.Id,
		)
		//log.Println(coin, "COIN", err)
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
			err = th.usersSpendsManager.CreateUserSpend(
				UsersSpendsManager.CreateUserSpendData{
					TxId:           tx.Id,
					TxSpendIndex:   spendIndex,
					UsersAddressId: userAddress.Id,
					Value:          spend.Value,
				},
			)
			if err != nil {
				continue
			}
			participateUserIdsMap[userAddress.UserId] = struct{}{}
		}
	}

	participateUserIds := make([]int, 0)
	for id := range participateUserIdsMap {
		participateUserIds = append(participateUserIds, id)
	}

	return tx, participateUserIds
}

func (th *TxsHandler) Start() {}
func (*TxsHandler) Status() error {
	return nil
}
func (*TxsHandler) Stop() error {
	return nil
}
