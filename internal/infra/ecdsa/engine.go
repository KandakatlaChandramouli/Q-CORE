package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"

	qcrypto "github.com/KandakatlaChandramouli/Q-CORE/internal/core/crypto"
)

var (
	ErrInvalidPrivateKey = errors.New("ecdsa: invalid private key")
	ErrInvalidPublicKey  = errors.New("ecdsa: invalid public key")
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (Engine) Algorithm() qcrypto.Algorithm {
	return qcrypto.AlgorithmECDSAP256
}

func (Engine) Metadata() qcrypto.Metadata {
	return qcrypto.Metadata{
		Algorithm:         qcrypto.AlgorithmECDSAP256,
		PublicKeyBytes:    65,
		SignatureBytes:    72,
		SecurityBits:      128,
		Standard:          "FIPS 186-5",
		VariableSignature: true,
		Deterministic:     false,
	}
}

func (Engine) GenerateKey() (qcrypto.PublicKey, qcrypto.PrivateKey, error) {
	sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	return &sk.PublicKey, sk, nil
}
