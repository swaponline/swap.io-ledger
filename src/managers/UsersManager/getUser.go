package UsersManager

import "swap.io-ledger/src/database"

func (um *UsersManager) GetUser(id int) (*database.User, error) {
	user, err := um.database.UsersGetById(id)

	return user, err
}
