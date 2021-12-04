package HttpServer

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"swap.io-ledger/src/auth"
	"swap.io-ledger/src/database"
)

func (hs *HttpServer) handleRegistration() {
	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		var registrationData RegistrationData
		if err := json.NewDecoder(r.Body).Decode(&registrationData); err != nil {
			log.Println("WARN:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid data"))
			return
		}

		ok := auth.VerifySign(
			registrationData.Addresses,
			registrationData.Sign,
			registrationData.PubKey,
		)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("signature invalid"))
			return
		}

		addressesBytes, err := hex.DecodeString(registrationData.Addresses)
		if err != nil {
			log.Println("WARN:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid data(addresses hex)"))
			return
		}
		var addresses []database.CreateUserAddressData
		if err := json.Unmarshal(addressesBytes, &addresses); err != nil {
			log.Println("WARN:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid data(addresses hex to json)"))
			return
		}

		err = hs.registrar.RegistrarUser(
			registrationData.PubKey,
			addresses,
		)
		if err != nil {
			log.Println("WARN:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("not create user in db"))
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
