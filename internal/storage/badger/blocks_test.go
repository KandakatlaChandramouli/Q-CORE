package badger

import (
	"context"
	"testing"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/block"
)

func TestStoreLoadBlock(t *testing.T) {
	engine := New(Options{
		InMemory: true,
	})

	if err := engine.Open(context.Background()); err != nil {
		t.Fatal(err)
	}
	defer engine.Close()

	b := block.New(7, [32]byte{}, nil)

	if err := engine.StoreBlock(context.Background(), &b); err != nil {
		t.Fatal(err)
	}

	out, err := engine.LoadBlock(context.Background(), 7)
	if err != nil {
		t.Fatal(err)
	}

	if out.Header.Height != 7 {
		t.Fatal("height mismatch")
	}
}
