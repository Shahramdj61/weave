package cash

import (
	"github.com/confio/weave"
	"github.com/confio/weave/errors"
)

const optKey = "cash"

// GenesisAccount is used to parse the json from genesis file
// use weave.Address, so address in hex, not base64
type GenesisAccount struct {
	Address weave.Address `json:"address"`
	Set
}

// Initializer fulfils the InitStater interface to load data from
// the genesis file
type Initializer struct{}

var _ weave.Initializer = Initializer{}

// FromGenesis will parse initial account info from genesis
// and save it to the database
func (Initializer) FromGenesis(opts weave.Options, kv weave.KVStore) error {
	accts := []GenesisAccount{}
	err := opts.ReadOptions(optKey, &accts)
	if err != nil {
		return err
	}
	bucket := NewBucket()
	for _, acct := range accts {
		// try to load up into a valid address
		if len(acct.Address) != weave.AddressLength {
			return errors.ErrUnrecognizedAddress(acct.Address)
		}
		wallet := NewWallet(acct.Address)
		err := wallet.Concat(acct.Set.Coins)
		if err != nil {
			return err
		}
		err = bucket.Save(kv, wallet)
		if err != nil {
			return err
		}
	}
	return nil
}
