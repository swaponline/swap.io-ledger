package txsHandler

type NonHandledTx struct {
	Hash              string       `json:"hash"`
	Journal           []SpendsInfo `json:"journal"`
	AllSpendAddresses []string
}
type SpendsInfo struct {
	Asset   SpendsAsset `json:"asset"`
	Entries []Spend     `json:"entries"`
}
type SpendsAsset struct {
	Id      string `json:"id"`
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Network string `json:"network"`
}
type Spend struct {
	Wallet string `json:"wallet,omitempty"`
	Value  string `json:"value"`
	Label  string `json:"label"`
}
