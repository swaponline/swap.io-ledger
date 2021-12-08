package registrar

import (
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
				// todo: subscribe in goroutine
				// todo: if error save subscribe and repeat feature
				if agentHandler, ok := (*r.networks)[address.Network]; ok {
					err := agentHandler.Subscribe(address.Address)
					return err
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
