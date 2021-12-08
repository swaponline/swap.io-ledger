package database

type Tx struct {
	Id   int
	Hash string
	Data string
}
type Network struct {
	Id   int
	Name string
}
type Coin struct {
	Id   int
	Name string
}
type User struct {
	Id     int
	PubKey string
}
type UserAddress struct {
	Id      int
	CoinId  int
	UserId  int
	Address string
}
type CreateUserAddressData struct {
	Network string
	Coin    string
	Address string
}
type AddressSyncStatus struct {
	AddressId int
	Sync      int
	Cursor    string
	Address   string
	Network   string
}
type SummaryUserSpends struct {
	Network string
	Coin    string
	value   string
}
