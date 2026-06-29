package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"

	qcrypto "github.com/KandakatlaChandramouli/Q-CORE/internal/core/crypto"
)

var ErrInvalidPrivateKey = errors.New("ecdsa: invalid private key")
var ErrInvalidPublicKey = errors.New("ecdsa: invalid public key")

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (Engine) Algorithm() qcrypto.Algorithm {
	return qcrypto.AlgorithmECDSAP256
}

func (Engine) GenerateKey() (qcrypto.PublicKey, qcrypto.PrivateKey, error) {
	sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return &sk.PublicKey, sk, nil
}

func (Engine) PublicKeySize() int {
	return 65
}

func (Engine) SignatureSize() int {
	return 72
}
