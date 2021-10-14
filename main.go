package main

import (
    "swap.io-ledger/src/config"
    "swap.io-ledger/src/agentHandler"
)

func main() {
	config.InitializeConfig()
    hsd := config.AGENTS[0]
    agentHandler.InitialiseAgentHandler(
        agentHandler.Config{
            Network: hsd.Network,
            BaseUrl: hsd.BaseUrl,
            ApiKey: hsd.ApiKey,
        },
    )

    <-make(chan struct{})
}
