package database

type Tx struct {
	Id int
	Hash string
	Data string
}
type Coin struct {
	Id int
	Name string
}
type User struct {
	Id int
	PubKey string
}
type UserAddress struct {
	Id int
	CoinId int
	UserId int
	Address string
}
type UserSpend struct {
	TxId int
	TxWiringIndex int
	UserAddressId int
	Value int
}