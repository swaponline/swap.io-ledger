package networks

import (
	"log"
	AgentHandler "swap.io-ledger/src/agentHandler"
	appConfig "swap.io-ledger/src/config"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/txsHandler"
)

type Networks map[string]*AgentHandler.AgentHandler
type Config struct {
	TxsHandler *txsHandler.TxsHandler
}

func InitialiseNetworks(config Config) *Networks {
	networks := make(Networks)

	for _, agentInfo := range appConfig.AGENTS {
		networks[agentInfo.Network] = AgentHandler.InitialiseAgentHandler(AgentHandler.Config{
			Network:    agentInfo.Network,
			ApiKey:     agentInfo.ApiKey,
			BaseUrl:    agentInfo.BaseUrl,
			TxsHandler: config.TxsHandler,
		})
	}

	return &networks
}
func Register(
	reg *serviceRegistry.ServiceRegistry,
) {
	var txsHandlerInstance *txsHandler.TxsHandler
	err := reg.FetchService(&txsHandlerInstance)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(InitialiseNetworks(Config{
		TxsHandler: txsHandlerInstance,
	}))
	if err != nil {
		log.Panicln(err)
	}
}

func (n *Networks) Start() {
	for _, agentHandler := range *n {
		go agentHandler.Start()
	}
}
func (n *Networks) Status() error {
	for _, agentHandler := range *n {
		err := agentHandler.Status()
		if err != nil {
			return err
		}
	}
	return nil
}
func (n *Networks) Stop() error {
	for _, agentHandler := range *n {
		err := agentHandler.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}
