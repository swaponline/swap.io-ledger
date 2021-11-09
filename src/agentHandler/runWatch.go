package agentHandler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"swap.io-ledger/src/txsHandler"
)

func (a *AgentHandler) runWatch() {
	u := url.URL{
		Scheme: "ws",
		Host: a.baseUrl,
		Path: "/ws",
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

		var aTx txsHandler.NonHandledTx
		if err := json.Unmarshal(msg,&aTx); err != nil {
			log.Println("ERROR:", err)
			continue
		}
		log.Println("on tx", aTx.Hash)

		tx := a.txsHandler.TxHandle(&aTx)

		err = c.WriteMessage(websocket.TextMessage, []byte{})
		if err != nil {
			log.Println("ERROR:", err)
			break
		}
		log.Println("tx receive", tx.Id, aTx.Hash)

		//a.TxsSource <- tx
	}
}
