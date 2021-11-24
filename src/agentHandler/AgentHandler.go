package AgentHandler

import (
	"log"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/txsHandler"
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

func (a *AgentHandler) Start() {
	for {
		a.runWatch()
	}
}
func (*AgentHandler) Status() error {
    return nil
}
func (*AgentHandler) Stop() error {
    return nil
}
