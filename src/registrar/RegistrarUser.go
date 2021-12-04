package registrar

import (
	"log"
	"swap.io-ledger/src/database"
)

func (r *Registrar) RegistrarUser(
	pubKey string,
	addresses []database.CreateUserAddressData,
) error {
	_, newAddressesIds, err := r.usersManager.CreateUserByPubKeyAndAddresses(
		pubKey,
		addresses,
		func() error {
			for _, address := range addresses {
				if address.Coin == "HSN" {
					log.Println("subscribe on", address.Address)
					return r.agentHandler.Subscribe(address.Address)
				}
			}
			return nil
		},
	)
	if err != nil {
		return err
	}

	r.addressSyncer.SyncNewAddresses(newAddressesIds)
	return nil
}
