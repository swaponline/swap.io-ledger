package main

import (
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/config"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/httpHandler"
	"swap.io-ledger/src/serviceRegistry"
	"swap.io-ledger/src/socketServer"
	"swap.io-ledger/src/txsHandler"
	"swap.io-ledger/src/usersManager"
)

func main() {
	config.InitializeConfig()

    registry := serviceRegistry.NewServiceRegistry()

    database := database.InitialiseDatabase()
    registry.RegisterService(
        database,
    )

    usersManager := usersManager.InitialiseUsersManager(
        usersManager.Config{},
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
            TxSource: agentHandler.TxsSource,
            TxIsReceive: agentHandler.TxIsReceive,
            Database: database,
            UsersManager: usersManager,
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
