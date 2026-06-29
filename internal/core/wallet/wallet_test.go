package wallet

import (
	"testing"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/infra/ecdsa"
)

func TestWallet(t *testing.T) {
	engine := ecdsa.New()

	w, err := New(engine)
	if err != nil {
		t.Fatal(err)
	}

	msg := []byte("hello q-core")

	sig, err := w.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}

	if err := engine.Verify(w.PublicKey(), msg, sig); err != nil {
		t.Fatal(err)
	}
}
