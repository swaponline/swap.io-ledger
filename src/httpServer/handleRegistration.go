package httpserver

import "net/http"

func handleRegistration() {
    http.HandleFunc(
        "/registration",
        func(w http.ResponseWriter, r *http.Request) {
        },
    )
}
