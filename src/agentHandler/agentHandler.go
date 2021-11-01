package agentHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type AgentHandler struct {
	network     string
	baseUrl     string
	apiKey      string
	TxsSource   chan *AgentTx
	TxIsReceive chan struct{}
	isSync      chan struct{}
}

type Config struct {
	Network string
	BaseUrl string
	ApiKey  string
}

func InitialiseAgentHandler(config Config) *AgentHandler {
	handler := AgentHandler{
		network:     config.Network,
		baseUrl:     config.BaseUrl,
		apiKey:      config.ApiKey,
		TxsSource:         make(chan *AgentTx),
		TxIsReceive: make(chan struct{}),
		isSync:      make(chan struct{}),
	}

    go func() {
        for {
            handler.run()
        }
    }()

	return &handler
}

func (a *AgentHandler) run() {
    u := url.URL{
        Scheme: "ws",
        Host: a.baseUrl,
        Path: "/ws",
        RawQuery: fmt.Sprintf("token=%v", a.apiKey),
    }
    c, _, err := websocket.DefaultDialer.Dial(
        u.String(),
        nil,
    )
    if err != nil {
        log.Panicln(err)
    }
    defer c.Close()
    log.Printf(
        "connected agent(network:%v|baseUrl:%v)",
        a.network,
        a.baseUrl,
    )

    connectIsDeactivate, deactivate := context.WithCancel(context.Background())
    go func() {
        for {
            select {
            case <- a.TxIsReceive:
                {
                    err := c.WriteMessage(websocket.TextMessage, []byte{})
                    if err != nil {
                        log.Println("ERROR:", err)
                        deactivate()
                        return
                    }
                }
            case <-connectIsDeactivate.Done():
                return
            }
        }
    }()

    for {
        _, msg, err := c.ReadMessage()
        if err != nil {
            log.Println("ERROR:", err)
            deactivate()
            return
        }

        var tx *AgentTx
        if err := json.Unmarshal(msg,&tx); err != nil {
            log.Println("ERROR:", err)
            continue
        }

        select {
        case a.TxsSource <- tx:
            continue
        case <-connectIsDeactivate.Done():
            return
        }
    }
}

func (*AgentHandler) Start() {}
func (*AgentHandler) Status() error {
    return nil
}
func (*AgentHandler) Stop() error {
    return nil
}
