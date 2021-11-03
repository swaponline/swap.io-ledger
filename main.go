package main

import (
    "swap.io-ledger/src/agentHandler"
    "swap.io-ledger/src/config"
    "swap.io-ledger/src/database"
    "swap.io-ledger/src/httpHandler"
    "swap.io-ledger/src/managers/UsersManager"
    "swap.io-ledger/src/serviceRegistry"
    "swap.io-ledger/src/socketServer"
    "swap.io-ledger/src/txsHandler"
)

func main() {
	config.InitializeConfig()

    registry := serviceRegistry.NewServiceRegistry()

    database := database.InitialiseDatabase()
    registry.RegisterService(
        database,
    )

    usersManager := UsersManager.InitialiseUsersManager(
        UsersManager.Config{},
    )
    registry.RegisterService(usersManager)

    hsd := config.AGENTS[0]
    agentHandler := agentHandler.InitialiseAgentHandler(
        agentHandler.Config{
            Network: hsd.Network,
            BaseUrl: hsd.BaseUrl,
            ApiKey: hsd.ApiKey,
        },
    )
    registry.RegisterService(
        agentHandler,
    )

    txsHandler := txsHandler.InitialiseTxsHandler(
        txsHandler.Config{
            Network: hsd.Network,
            ATxSource: agentHandler.TxsSource,
            ATxIsReceive: agentHandler.TxIsReceive,
            //todo: add handlers
        },
    )
    registry.RegisterService(
        txsHandler,
    )

    registry.RegisterService(
        socketServer.InitialiseSocketServer(),
    )
    registry.RegisterService(
        httpHandler.InitializeServer(),
    )

    registry.StartAll()

    <-make(chan struct{})
}
