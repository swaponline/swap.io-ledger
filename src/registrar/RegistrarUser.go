package registrar

import "swap.io-ledger/src/database"

func (r *Registrar) RegistrarUser(
	pubKey string,
	addresses []database.CreateUserAddressData,
) error {
	_, newAddressesIds, err := r.usersManager.CreateUserByPubKeyAndAddresses(
		pubKey, addresses,
	)
	if err != nil {
		return err
	}

	r.addressSyncer.SyncNewAddresses(newAddressesIds)
	return nil
};