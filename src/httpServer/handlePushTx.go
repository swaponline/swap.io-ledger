package HttpServer

import "net/http"

func (hs *HttpServer) handlePushTx() {
	http.HandleFunc("/pushTx", func(w http.ResponseWriter, r *http.Request) {

	})
}
