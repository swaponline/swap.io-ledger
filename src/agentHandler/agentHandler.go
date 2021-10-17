package agentHandler

import (
	"log"
    "net/url"
    "encoding/json"
    "context"
    "fmt"

    "github.com/gorilla/websocket"
)

type AgentHandler struct {
	network     string
	baseUrl     string
	apiKey      string
	txs         chan<- *Transaction
	txIsReceive <-chan struct{}
	isSync      <-chan struct{}
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
		txs:         make(chan<- *Transaction),
		txIsReceive: make(<-chan struct{}),
		isSync:      make(<-chan struct{}),
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

    connectIsDeactive, deactive := context.WithCancel(context.Background())
    go func() {
        for {
            select {
            case <- a.txIsReceive:
                {
                    err := c.WriteMessage(websocket.TextMessage, []byte{})
                    if err != nil {
                        log.Println("ERROR:", err)
                        deactive()
                        return
                    }
                }
            case <-connectIsDeactive.Done():
                return
            }
        }
    }()

    for {
        _, msg, err := c.ReadMessage()
        if err != nil {
            log.Println("ERROR:", err)
            deactive()
            return
        }

        var tx *Transaction
        if err := json.Unmarshal(msg,&tx); err != nil {
            log.Println("ERROR:", err)
            continue
        }

        select {
        case a.txs <- tx:
            continue
        case <-connectIsDeactive.Done():
            return
        }
    }
}
