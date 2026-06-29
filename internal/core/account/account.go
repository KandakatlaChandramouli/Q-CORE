package account

import "github.com/KandakatlaChandramouli/Q-CORE/internal/core/address"

type Account struct {
	Address address.Address

	Balance uint64

	Nonce uint64
}
