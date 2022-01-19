package HttpServer

import (
	"encoding/json"
	"log"
	"net/http"
	"swap.io-ledger/src/utils"
)

func (hs *HttpServer) handleGetTxs() {
	http.HandleFunc("/getTxs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		userId, err := hs.auth.AuthenticationRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userTxs, err := hs.txsManager.GetTxsByUserId(userId)
		if err != nil {
			log.Println("ERROR:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userTxsBytes, err := json.Marshal(userTxs)
		if err != nil {
			log.Println("ERROR:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		utils.SetContentTypeJson(w)
		w.WriteHeader(http.StatusOK)
		w.Write(userTxsBytes)
	})
}
