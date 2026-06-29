package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"

	qcrypto "github.com/KandakatlaChandramouli/Q-CORE/internal/core/crypto"
)

const Version uint16 = 1

type Transaction struct {
	Version uint16

	ID [32]byte

	From qcrypto.PublicKey
	To   qcrypto.PublicKey

	Amount uint64
	Nonce  uint64

	Signature []byte
}

func (tx *Transaction) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(tx); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (tx *Transaction) ComputeID() error {
	data, err := tx.Serialize()
	if err != nil {
		return err
	}

	sum := sha256.Sum256(data)

	copy(tx.ID[:], sum[:])

	return nil
}

func New() Transaction {
	return Transaction{
		Version: Version,
	}
}

func init() {
	gob.Register([]byte{})
	gob.Register([32]byte{})
	binary.BigEndian.Uint64(make([]byte, 8))
}
