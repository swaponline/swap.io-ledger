package HttpServer

import (
	"encoding/json"
	"log"
	"net/http"
)

func (hs *HttpServer) handleGetBalances() {
	http.HandleFunc("/getBalances", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		userId, err := hs.auth.AuthenticationRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		summaryUserSpends, err := hs.usersSpendsManager.GetSummarySpends(userId)
		if err != nil {
			log.Panicln("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		summaryUserSpendsBytes, err := json.Marshal(summaryUserSpends)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(summaryUserSpendsBytes)
	})
}
