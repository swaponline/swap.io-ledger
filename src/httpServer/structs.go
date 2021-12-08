package HttpServer

type RegistrationData struct {
	PubKey    string `json:"pubKey"`
	Addresses string `json:"addresses"`
	Sign      string `json:"sign"`
}
type PushTxData struct {
	Network string `json:"network"`
	Hex     string `json:"hex"`
}
