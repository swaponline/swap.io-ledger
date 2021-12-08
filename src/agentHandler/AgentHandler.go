package AgentHandler

import (
	"swap.io-ledger/src/txsHandler"
)

type AgentHandler struct {
	network         string
	baseUrl         string
	apiKey          string
	txsHandler      *txsHandler.TxsHandler
	TxNotifications chan *TxNotification
}

type Config struct {
	Network    string
	BaseUrl    string
	ApiKey     string
	TxsHandler *txsHandler.TxsHandler
}

func InitialiseAgentHandler(config Config) *AgentHandler {
	handler := AgentHandler{
		network:         config.Network,
		baseUrl:         config.BaseUrl,
		apiKey:          config.ApiKey,
		txsHandler:      config.TxsHandler,
		TxNotifications: make(chan *TxNotification),
	}

	return &handler
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
