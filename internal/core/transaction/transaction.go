package transaction

import (
	"crypto/sha256"
	"encoding/binary"

	qcrypto "github.com/KandakatlaChandramouli/Q-CORE/internal/core/crypto"
)

type Transaction struct {
	ID        [32]byte
	From      qcrypto.PublicKey
	To        qcrypto.PublicKey
	Amount    uint64
	Nonce     uint64
	Signature []byte
}

func (tx *Transaction) ComputeID() {
	h := sha256.New()

	var buf [16]byte

	binary.BigEndian.PutUint64(buf[:8], tx.Amount)
	binary.BigEndian.PutUint64(buf[8:], tx.Nonce)

	h.Write(buf[:])

	if tx.Signature != nil {
		h.Write(tx.Signature)
	}

	copy(tx.ID[:], h.Sum(nil))
}
