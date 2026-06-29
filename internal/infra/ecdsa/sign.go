package ecdsa

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"

	qcrypto "github.com/KandakatlaChandramouli/Q-CORE/internal/core/crypto"
)

func (Engine) Sign(
	privateKey qcrypto.PrivateKey,
	message []byte,
) ([]byte, error) {

	sk, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, ErrInvalidPrivateKey
	}

	hash := sha256.Sum256(message)

	return ecdsa.SignASN1(rand.Reader, sk, hash[:])
}

func (Engine) Verify(
	publicKey qcrypto.PublicKey,
	message []byte,
	signature []byte,
) error {

	pk, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return ErrInvalidPublicKey
	}

	hash := sha256.Sum256(message)

	if !ecdsa.VerifyASN1(pk, hash[:], signature) {
		return errors.New("ecdsa: invalid signature")
	}

	return nil
}
