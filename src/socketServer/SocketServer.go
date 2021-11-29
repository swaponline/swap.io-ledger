package socketServer

import (
	"encoding/json"
	"log"
	"net/http"
	"swap.io-ledger/src/database"
	"swap.io-ledger/src/serviceRegistry"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/auth"
)

type SocketServer struct {
	auth *auth.Auth
    txSource <-chan *AgentHandler.TxNotification
}

type Config struct {
	Auth *auth.Auth
    agentHandlers []*AgentHandler.AgentHandler
	txSource <- chan AgentHandler.TxNotification
}

const writePeriod = time.Minute * 1
const readPeriod  = time.Minute * 2

var upgrader = websocket.Upgrader{}

func InitialiseSocketServer(config Config) *SocketServer {
	type UserListener struct {
		pingTicker *time.Ticker
		txNotification chan *database.Tx
	}
	userListeners := make(map[int]UserListener)
	userListenersLocker := sync.Mutex{}
	go func() {
		for {
			userListenersLocker.Lock()
			txNotification := <- config.txSource
			for _, userId := range txNotification.UsersIds {
				if userListener, ok := userListeners[userId]; ok {
					userListener.txNotification <- txNotification.Tx
				}
			}
			userListenersLocker.Unlock()
		}
	}()

    wsHandle := func(w http.ResponseWriter, r *http.Request) {
        userId, err := config.Auth.AuthenticationRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
			return
		}

        c, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Println("upgrade:", err)
            return
        }
        defer c.Close()

		ticker := time.NewTicker(writePeriod)
		defer ticker.Stop()
		txNotification := make(chan *database.Tx)
		defer close(txNotification)

		userListener := UserListener{
			pingTicker: ticker,
			txNotification: txNotification,
		}
		userListeners[userId] = userListener

        log.Println("connect:", userId)

		go func() {
            for {
				select {
                case <-userListener.pingTicker.C:
                    {
                        log.Println("ping")
                        c.SetWriteDeadline(time.Now().Add(writePeriod))
                        if err := c.WriteMessage(
                            websocket.PingMessage, nil,
                        ); err != nil {
                            log.Println("ERROR (ticker)", err)
                            return
                        }
                    }
				case tx, ok := <-userListener.txNotification:
					{
						if sendingData, err := json.Marshal(tx); err == nil && ok {
							if err := c.WriteMessage(
								websocket.TextMessage,
								sendingData,
							); err != nil {
								return
							}
						}
					}
				}
			}
		}()

        c.SetReadDeadline(time.Now().Add(readPeriod))
        c.SetPongHandler(
            func(string) error {
                log.Println("pong")
                c.SetReadDeadline(time.Now().Add(readPeriod))
                return nil
            },
        )
        for {
            _, _, err := c.ReadMessage()
            if err != nil {
                log.Println("disconnect:", userId)
                return
            }
        }

		userListenersLocker.Lock()
		delete(userListeners, userId)
		userListenersLocker.Unlock()
    }

    http.HandleFunc("/ws", wsHandle)

	return &SocketServer{}
}

func Register(reg *serviceRegistry.ServiceRegistry) {
	var auth *auth.Auth
	err := reg.FetchService(&auth)
	if err != nil {
		log.Panicln(err)
	}

	err = reg.RegisterService(InitialiseSocketServer(Config{
		Auth: auth,
	}))
	if err != nil {
		log.Panicln(err)
	}
}

func (*SocketServer) Start() {}
func (*SocketServer) Status() error {
    return nil
}
func (*SocketServer) Stop() error {
    return nil
}
