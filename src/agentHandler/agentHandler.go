package agentHandler

import (
	"encoding/json"
	"fmt"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
	"log"
)

type AgentHandler struct {
	network     string
	baseUrl     string
	apiKey      string
	txs         chan <- *Transaction
	txIsReceive <- chan struct{}
	isSync <-chan struct{}
}

type Config struct {
	Network string
	BaseUrl string
	ApiKey string
}

func initialiseAgentHandler(config Config) *AgentHandler {
	handler := AgentHandler{
		network: config.Network,
		baseUrl: config.BaseUrl,
		apiKey:  config.ApiKey,
		txs: make(chan <- *Transaction),
		txIsReceive: make(<- chan struct{}),
		isSync: make(<-chan struct{}),
	}

	go handler.run()

	return &handler
}

func (a *AgentHandler) run() {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["token"] = a.apiKey
	uri := fmt.Sprintf("%v/socket.io/", a.baseUrl)

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Panicf("NewClient error:%v", err)
	}

	err = client.On("error", func() {
		log.Println("on error")
	})
	if err != nil {
		log.Panicln(err)
	}

	err = client.On("connection", func() {
		log.Println("on connect")
	})
	if err != nil {
		log.Panicln(err)
	}

	err = client.On("newTransaction", func(txStr string) {
		var tx *Transaction
		if err = json.Unmarshal([]byte(txStr), &tx); err != nil {
			log.Println("ERROR", err, txStr)
		}
		a.txs <- tx
		<- a.txIsReceive
		log.Printf("%v| tx receive :%v", a.network, tx.Hash)
	})
	if err != nil {
		log.Panicln(err)
	}

	err = client.On("disconnection", func() {
		log.Println("on disconnect")
	})
	if err != nil {
		log.Panicln(err)
	}
}