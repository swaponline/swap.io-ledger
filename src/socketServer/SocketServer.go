package socketServer

import (
	"log"
	"net/http"
	"swap.io-ledger/src/serviceRegistry"
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
}

const writePeriod = time.Minute * 1
const readPeriod  = time.Minute * 2

var upgrader = websocket.Upgrader{}

func InitialiseSocketServer(config Config) *SocketServer {
    wsHandle := func(w http.ResponseWriter, r *http.Request) {
        userId, err := config.Auth.AuthenticationRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
			return
		}
        //if err != nil {
        //    log.Println("ERROR user not connected")
        //    w.WriteHeader(http.StatusUnauthorized)
        //    w.Write([]byte(`failed auth`))
        //    return
        //}

        c, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Println("upgrade:", err)
            return
        }
        defer c.Close()

        log.Println("connect:", userId)

        ticker := time.NewTicker(writePeriod)
        defer ticker.Stop()

		go func() {
            for {
				select {
                case <-ticker.C:
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
