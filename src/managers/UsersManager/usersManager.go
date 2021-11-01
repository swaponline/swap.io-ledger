package UsersManager

import "swap.io-ledger/src/database"

type UsersManager struct {
    database *database.Database
}
type Config struct {
    Database *database.Database
}

func InitialiseUsersManager(config Config) *UsersManager {
    return &UsersManager{
        database: config.Database,
    }
}

func (*UsersManager) Start() {}
func (*UsersManager) Status() error {
    return nil
}
func (*UsersManager) Stop() error {
    return nil
}
