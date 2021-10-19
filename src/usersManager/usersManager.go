package usersManager

type UsersManager struct {}
type Config struct {}

func InitialiseUsersManager(config Config) *UsersManager {
    return &UsersManager{}
}

func (*UsersManager) Start() {}
func (*UsersManager) Status() error {
    return nil
}
func (*UsersManager) Stop() error {
    return nil
}
