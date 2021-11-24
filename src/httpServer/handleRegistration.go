package Httpserver

import (
    "encoding/json"
    "net/http"
    "swap.io-ledger/src/auth"
)

func (hs *HttpServer) handleRegistration() {
    http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
        var registrationData RegistrationData
        if err := json.NewDecoder(r.Body).Decode(&registrationData); err != nil {
            w.Write([]byte("invalid data"))
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        err := auth.VerifySign(
            registrationData.Addresses,
            registrationData.Sign,
            registrationData.PubKey,
        )
        if err {
            w.Write([]byte("signature invalid"))
            w.WriteHeader(http.StatusBadRequest)
            return
        }
    })
}
