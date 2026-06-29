package blockchain

import (
	"testing"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/block"
)

func TestBlockchain(t *testing.T) {
	bc := New()

	if bc.Height() != 0 {
		t.Fatal("expected empty chain")
	}

	b := block.New(1, [32]byte{}, nil)

	bc.AddBlock(b)

	if bc.Height() != 1 {
		t.Fatal("unexpected height")
	}

	_, ok := bc.Latest()
	if !ok {
		t.Fatal("missing latest block")
	}
}
