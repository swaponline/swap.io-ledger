package UsersAdressesManager

import "swap.io-ledger/src/database"

func (ua *UsersAddressesManager) GetUserAddress(
	coinId int,
	address string,
) (*database.UserAddress, error) {
	return ua.database.UsersAddressesGetByCoinIdAndAddress(
		coinId,
		address,
	)
}
