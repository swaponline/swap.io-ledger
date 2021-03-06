package database

import shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"

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
type UserBalance struct {
	Network string             `json:"network"`
	Coin    string             `json:"coin"`
	Balance shopspring.Numeric `json:"balance"`
}
