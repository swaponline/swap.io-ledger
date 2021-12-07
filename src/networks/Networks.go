package networks

import (
	AgentHandler "swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/config"
	"swap.io-ledger/src/txsHandler"
)

type Networks map[string]*AgentHandler.AgentHandler
type Config struct {
	TxsHandler *txsHandler.TxsHandler
}

func InitialiseNetworks() *Networks {
	networks := new(Networks)

	for _, agent := range config.AGENTS {
		(*networks)[agent.Network] = AgentHandler.InitialiseAgentHandler(AgentHandler.Config{
			Network: agent.Network,
			ApiKey:  agent.ApiKey,
			BaseUrl: agent.BaseUrl,
		})
	}

	return networks
}

func (*Networks) Start() {}
func (*Networks) Status() error {
	return nil
}
func (*Networks) Stop() error {
	return nil
}
