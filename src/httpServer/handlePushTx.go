package HttpServer

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
)

func (hs *HttpServer) handlePushTx() {
	http.HandleFunc("/pushTx", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var pushTxData PushTxData
		if err := json.NewDecoder(r.Body).Decode(&pushTxData); err != nil {
			log.Println("WARN:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("not decode payload in struct"))
			return
		}

		agentHandler, ok := (*hs.networks)[pushTxData.Network]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not exist emitted network"))
			return
		}

		hexBytes, err := hex.DecodeString(pushTxData.Hex)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid hex"))
			return
		}

		agentResponseBytes, err := agentHandler.PushTx(hexBytes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(agentResponseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(agentResponseBytes)
	})
}
