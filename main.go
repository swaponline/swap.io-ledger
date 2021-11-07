package main

import (
    "log"
    "swap.io-ledger/src/agentHandler"
    "swap.io-ledger/src/config"
    "swap.io-ledger/src/database"
    "swap.io-ledger/src/httpHandler"
    "swap.io-ledger/src/managers/CoinsManager"
    "swap.io-ledger/src/managers/TxsManager"
    "swap.io-ledger/src/managers/UsersAdressesManager"
    "swap.io-ledger/src/managers/UsersManager"
    "swap.io-ledger/src/managers/UsersSpendsManager"
    "swap.io-ledger/src/serviceRegistry"
    "swap.io-ledger/src/socketServer"
    "swap.io-ledger/src/txsHandler"
)

func main() {
	config.InitializeConfig()

    registry := serviceRegistry.NewServiceRegistry()

    databaseInstance := database.InitialiseDatabase()
    err := registry.RegisterService(
        databaseInstance,
    )
    if err != nil {
        log.Panicln(err)
    }

    TxsManager.Register(registry)
    CoinsManager.Register(registry)
    UsersManager.Register(registry)
    UsersAdressesManager.Register(registry)
    UsersSpendsManager.Register(registry)

    txsHandler.Register(registry)

    hsd := config.AGENTS[0]
    err = agentHandler.Register(
        registry,
        hsd.Network,
        hsd.BaseUrl,
        hsd.ApiKey,
    )
    if err != nil {
        log.Panicln(err)
    }

    err = registry.RegisterService(
        socketServer.InitialiseSocketServer(),
    )
    if err != nil {
        log.Panicln(err)
    }

    err = registry.RegisterService(
        httpHandler.InitializeServer(),
    )
    if err != nil {
        log.Panicln(err)
    }

    registry.StartAll()

    <-make(chan struct{})
}
