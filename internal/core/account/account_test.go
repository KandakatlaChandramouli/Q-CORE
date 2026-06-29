package account

import "testing"

func TestAccountDefaults(t *testing.T) {
	var acc Account

	if acc.Balance != 0 {
		t.Fatal("balance must start at zero")
	}

	if acc.Nonce != 0 {
		t.Fatal("nonce must start at zero")
	}
}
