package UsersAdressesManager

import "swap.io-ledger/src/database"

func (ua *UsersAddressesManager) GetUserAddress(
	address string,
) (*database.UserAddress, error) {
	return ua.database.UsersAddressesGetByAddress(address)
}
