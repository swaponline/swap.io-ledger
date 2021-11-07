package agentHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/txsHandler"

	"github.com/gorilla/websocket"
)

type AgentHandler struct {
	network     string
	baseUrl     string
	apiKey      string
	txsHandler  *txsHandler.TxsHandler
	TxsSource   chan *database.Tx
}

type Config struct {
	Network string
	BaseUrl string
	ApiKey  string
	TxsHandler *txsHandler.TxsHandler
}

func InitialiseAgentHandler(config Config) *AgentHandler {
	handler := AgentHandler{
		network:     config.Network,
		baseUrl:     config.BaseUrl,
		apiKey:      config.ApiKey,
		txsHandler:  config.TxsHandler,
		TxsSource:   make(chan *database.Tx),
	}

	return &handler
}
func Register(
	reg *serviceRegistry.ServiceRegistry,
	network string,
	baseUrl string,
	apiKey string,
) error {
	var txsHandlerInstance *txsHandler.TxsHandler
	err := reg.FetchService(&txsHandlerInstance)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(
		InitialiseAgentHandler(Config{
			Network: network,
			BaseUrl: baseUrl,
			ApiKey: apiKey,
			TxsHandler: txsHandlerInstance,
		}),
	)

	return err
}

func (a *AgentHandler) run() {
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

func (a *AgentHandler) Start() {
	for {
		a.run()
	}
}
func (*AgentHandler) Status() error {
    return nil
}
func (*AgentHandler) Stop() error {
    return nil
}
