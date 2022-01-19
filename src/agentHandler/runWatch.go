package AgentHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"swap.io-ledger/src/txsHandler"
)

func (a *AgentHandler) runWatch() {
	u := url.URL{
		Scheme:   "ws",
		Host:     a.baseUrl,
		Path:     "/ws",
		RawQuery: fmt.Sprintf("token=%v", a.apiKey),
	}
	c, _, err := websocket.DefaultDialer.Dial(
		u.String(),
		nil,
	)
	if err != nil {
		log.Panicln(err)
	}
	defer c.Close()

	log.Printf(
		"connected agent(network:%v|baseUrl:%v)",
		a.network,
		a.baseUrl,
	)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("ERROR:", err)
			return
		}

		var nonTx txsHandler.NonHandledTx
		if err := json.Unmarshal(msg, &nonTx); err != nil {
			log.Println("ERROR:", err)
			continue
		}
		log.Println("on tx", nonTx.Hash)

		tx, participateUserIdsMap := a.txsHandler.TxHandle(&nonTx)

		err = c.WriteMessage(websocket.TextMessage, []byte{})
		if err != nil {
			log.Println("ERROR:", err)
			break
		}

		a.TxNotifications <- &TxNotification{
			Tx:       tx,
			UsersIds: participateUserIdsMap,
		}
		log.Println("tx receive", tx.Id, nonTx.Hash, participateUserIdsMap)
	}
}
