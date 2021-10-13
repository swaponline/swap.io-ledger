package agentHandler

import (
	"log"
    "net/url"
    "encoding/json"

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
    defer close(handler.txs)

    go func() {
        for {
            handler.run()
        }
    }()

	return &handler
}

func (a *AgentHandler) run() {
    u := url.URL{Scheme: "ws", Host: a.baseUrl, Path: "/ws"}
    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Panicln(err)
    }
    defer c.Close()
    log.Printf(
        "connected agent(network:%v|baseUrl:%v)",
        a.network,
        a.baseUrl,
    )

    //c.SetPingHandler(func(data string) error {
    //    return nil
    //})

    done := make(chan struct{})
    defer close(done)

    go func() {
        loop: for {
            select {
            case <- a.txIsReceive: {
                err := c.WriteMessage(websocket.TextMessage, []byte{})
                if err != nil {
                    log.Println("ERROR:", err)
                }
            }
            default: {
                break loop
            }
            }
        }
    }()

    loop: for {
        _, msg, err := c.ReadMessage()
        if err != nil {
            log.Println("ERROR:", err)
        }

        var tx *Transaction
        if err := json.Unmarshal(msg,&tx); err != nil {
            log.Println("ERROR:", err)
            continue
        }

        select {
        case a.txs <- tx: continue loop
        default: break loop
        }
    }
}
