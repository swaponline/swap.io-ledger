package UsersManager

import "swap.io-ledger/src/database"

func (um *UsersManager) GetUser(id int) (*database.User, error) {
	return um.database.UsersGetById(id)
}
func (um *UsersManager) GetUserByPubKey(pubKey string) (*database.User, error) {
	return um.database.UsersGetByPubKey(pubKey)
}