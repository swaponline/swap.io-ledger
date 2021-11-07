package UsersManager

import (
    "log"
    "swap.io-ledger/src/database"
    "swap.io-ledger/src/serviceRegistry"
)

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
func Register(reg *serviceRegistry.ServiceRegistry) {
    var database *database.Database
    err := reg.FetchService(&database)
    if err != nil {
        log.Panicln(err)
    }

    err = reg.RegisterService(
        InitialiseUsersManager(Config{
            Database: database,
        }),
    )
}

func (*UsersManager) Start() {}
func (*UsersManager) Status() error {
    return nil
}
func (*UsersManager) Stop() error {
    return nil
}
