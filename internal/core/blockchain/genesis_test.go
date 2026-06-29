package blockchain

import "testing"

func TestGenesis(t *testing.T) {
	bc := New()

	bc.InitGenesis()

	if bc.Height() != 1 {
		t.Fatalf("expected genesis block, got %d", bc.Height())
	}

	b, ok := bc.Latest()
	if !ok {
		t.Fatal("missing genesis")
	}

	if b.Header.Height != 0 {
		t.Fatal("genesis height must be zero")
	}
}
