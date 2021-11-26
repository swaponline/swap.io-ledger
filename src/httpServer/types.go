package HttpServer

type RegistrationData struct {
	PubKey string `json:"pubKey"`
	Addresses string `json:"addresses"`
	Sign string `json:"sign"`
}