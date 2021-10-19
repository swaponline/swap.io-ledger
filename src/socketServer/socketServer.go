package socketServer

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"swap.io-ledger/src/agentHandler"
	"swap.io-ledger/src/auth"
)

type SocketServer struct {
    txSource <-chan *agentHandler.Transaction
}

type Config struct {
    agentHandlers []*agentHandler.AgentHandler
}

const writePeriod = time.Minute * 1
const readPeriod  = time.Minute * 2

var upgrader = websocket.Upgrader{}

func InitialiseSocketServer() *SocketServer {
    wsHandle := func(w http.ResponseWriter, r *http.Request) {
        userId, err := auth.AuthenticationRequest(r)
        if err != nil {
            log.Println("ERROR user not connected")
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte(`failed auth`))
            return
        }

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

func (*SocketServer) Start() {}
func (*SocketServer) Status() error {
    return nil
}
func (*SocketServer) Stop() error {
    return nil
}
