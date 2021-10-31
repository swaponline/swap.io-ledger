package UsersManager

func (um *UsersManager) CreateUser(
	pubKey string,
) int {
	um.database.UsersCreate(pubKey)

	return 0
}