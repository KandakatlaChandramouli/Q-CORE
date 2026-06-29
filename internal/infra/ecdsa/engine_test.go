package ecdsa

import "testing"

func TestSignVerify(t *testing.T) {

	engine := New()

	pub, priv, err := engine.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	msg := []byte("Q-CORE")

	sig, err := engine.Sign(priv, msg)
	if err != nil {
		t.Fatal(err)
	}

	if err := engine.Verify(pub, msg, sig); err != nil {
		t.Fatal(err)
	}
}
