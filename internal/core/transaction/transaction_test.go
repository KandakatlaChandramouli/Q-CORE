package transaction

import "testing"

func TestComputeID(t *testing.T) {
	tx := Transaction{
		Amount: 100,
		Nonce:  1,
	}

	tx.ComputeID()

	var zero [32]byte

	if tx.ID == zero {
		t.Fatal("transaction id not computed")
	}
}
