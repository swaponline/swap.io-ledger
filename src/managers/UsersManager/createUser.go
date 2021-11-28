package UsersManager

import "swap.io-ledger/src/database"

func (um *UsersManager) CreateUser(
	pubKey string,
) (int, error) {
	return um.database.UsersCreate(pubKey)
}
func (um *UsersManager) CreateUserByPubKeyAndAddresses(
	pubKey string,
	addresses []database.CreateUserAddressData,
	beforeCommit func() error,
) (int, []int, error) {
	return um.database.UsersCreateByPubKeyAndAddresses(
		pubKey,
		addresses,
		beforeCommit,
	)
}