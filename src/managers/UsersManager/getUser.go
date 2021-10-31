package UsersManager

func (um *UsersManager) GetUser(id int) (*User, error) {
	user, err := um.database.UsersGetById(id)

	return user, err
}
