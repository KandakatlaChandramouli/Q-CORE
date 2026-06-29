package address

import "testing"

func TestAddressString(t *testing.T) {
	var a Address

	if len(a.String()) != 64 {
		t.Fatal("unexpected address length")
	}
}
