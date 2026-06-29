package badger

import (
	"context"
	"testing"
)

func TestOpenClose(t *testing.T) {
	engine := New(Options{
		InMemory: true,
	})

	if err := engine.Open(context.Background()); err != nil {
		t.Fatal(err)
	}

	if engine.db == nil {
		t.Fatal("database not opened")
	}

	if err := engine.Close(); err != nil {
		t.Fatal(err)
	}

	if engine.db != nil {
		t.Fatal("database not closed")
	}
}
