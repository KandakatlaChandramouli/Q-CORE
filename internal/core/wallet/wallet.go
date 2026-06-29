package wallet

import (
	qcrypto "github.com/KandakatlaChandramouli/Q-CORE/internal/core/crypto"
)

type Wallet struct {
	engine qcrypto.Engine

	publicKey  qcrypto.PublicKey
	privateKey qcrypto.PrivateKey
}

func New(engine qcrypto.Engine) (*Wallet, error) {
	pub, priv, err := engine.GenerateKey()
	if err != nil {
		return nil, err
	}

	return &Wallet{
		engine:     engine,
		publicKey:  pub,
		privateKey: priv,
	}, nil
}

func (w *Wallet) PublicKey() qcrypto.PublicKey {
	return w.publicKey
}

func (w *Wallet) Sign(msg []byte) ([]byte, error) {
	return w.engine.Sign(w.privateKey, msg)
}
