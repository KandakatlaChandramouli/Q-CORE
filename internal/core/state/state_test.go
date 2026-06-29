package state

import (
	"testing"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/account"
	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/address"
)

func TestCreateAndGet(t *testing.T) {
	st := New()

	var addr address.Address

	st.Create(account.Account{
		Address: addr,
		Balance: 100,
	})

	acc, ok := st.Get(addr)
	if !ok {
		t.Fatal("account not found")
	}

	if acc.Balance != 100 {
		t.Fatal("unexpected balance")
	}
}
