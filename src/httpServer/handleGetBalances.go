package HttpServer

import (
	"encoding/json"
	"log"
	"net/http"
	"swap.io-ledger/src/utils"
)

func (hs *HttpServer) handleGetBalances() {
	http.HandleFunc("/getBalances", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		userId, err := hs.auth.AuthenticationRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userBalances, err := hs.usersSpendsManager.GetUserBalances(userId)
		if err != nil {
			log.Panicln("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userBalancesBytes, err := json.Marshal(userBalances)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		utils.SetContentTypeJson(w)
		w.WriteHeader(http.StatusOK)
		w.Write(userBalancesBytes)
	})
}
