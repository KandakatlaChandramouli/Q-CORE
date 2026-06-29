package block

import "testing"

func TestNewHeader(t *testing.T) {
	h := NewHeader(1, [32]byte{}, [32]byte{})

	if h.Version != 1 {
		t.Fatal("unexpected version")
	}

	if h.Height != 1 {
		t.Fatal("unexpected height")
	}

	if h.Timestamp == 0 {
		t.Fatal("timestamp not initialized")
	}
}
