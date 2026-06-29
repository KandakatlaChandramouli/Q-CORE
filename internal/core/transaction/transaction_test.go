package transaction

import "testing"

func TestSerializeAndComputeID(t *testing.T) {
	tx := New()

	tx.Amount = 100
	tx.Nonce = 1

	data, err := tx.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 0 {
		t.Fatal("empty serialization")
	}

	if err := tx.ComputeID(); err != nil {
		t.Fatal(err)
	}

	var zero [32]byte

	if tx.ID == zero {
		t.Fatal("id not computed")
	}
}
