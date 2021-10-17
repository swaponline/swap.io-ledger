package socketServer

import (
    "swap.io-ledger/src/agentHandler"
)

type SocketServer struct {
    txSource <-chan *agentHandler.Transaction
}

type Config struct {
    agentHandlers []*agentHandler.AgentHandler
}

func InitialiseSocketServer() {
    

}
