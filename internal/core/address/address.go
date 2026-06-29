package address

import (
	"crypto/sha256"
	"encoding/hex"

	qcrypto "github.com/KandakatlaChandramouli/Q-CORE/internal/core/crypto"
)

const Size = 32

type Address [Size]byte

func FromPublicKey(pk qcrypto.PublicKey) Address {
	raw, ok := pk.(interface {
		Bytes() []byte
	})

	if !ok {
		panic("public key does not expose Bytes()")
	}

	sum := sha256.Sum256(raw.Bytes())

	var addr Address
	copy(addr[:], sum[:])

	return addr
}

func (a Address) Bytes() []byte {
	out := make([]byte, Size)
	copy(out, a[:])
	return out
}

func (a Address) String() string {
	return hex.EncodeToString(a[:])
}
