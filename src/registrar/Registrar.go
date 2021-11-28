package registrar

import (
	"log"
	"swap.io-ledger/src/addressSyncer"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/managers/UsersManager"
	"swap.io-ledger/src/serviceRegistry"
)

type Registrar struct {
	usersManager *UsersManager.UsersManager
	addressSyncer *AddressSyncer.AddressSyncer
	agentHandler *AgentHandler.AgentHandler
}
type Config struct {
	UsersManager *UsersManager.UsersManager
	AddressSyncer *AddressSyncer.AddressSyncer
	AgentHandler *AgentHandler.AgentHandler
}

func InitialiseRegistrar(config Config) *Registrar {
	return &Registrar{
		usersManager: config.UsersManager,
		addressSyncer: config.AddressSyncer,
		agentHandler: config.AgentHandler,
	}
}

func Register(reg *serviceRegistry.ServiceRegistry) {
	var usersManager *UsersManager.UsersManager
	err := reg.FetchService(&usersManager)
	if err != nil {
		log.Panicln(err)
	}

	var addressSyncer *AddressSyncer.AddressSyncer
	err = reg.FetchService(&addressSyncer)
	if err != nil {
		log.Panicln(err)
	}

	var agentHandler *AgentHandler.AgentHandler
	err = reg.FetchService(&agentHandler)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(
		InitialiseRegistrar(Config{
			UsersManager: usersManager,
			AddressSyncer: addressSyncer,
			AgentHandler: agentHandler,
		}),
	)
}

func (*Registrar) Start() {}
func (*Registrar) Stop() error {
	return nil
}
func (*Registrar) Status() error {
	return nil
}