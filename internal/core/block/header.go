package block

import "time"

type Header struct {
	Version uint16

	Height uint64

	Timestamp int64

	PreviousHash [32]byte

	MerkleRoot [32]byte

	Nonce uint64
}

func NewHeader(height uint64, previousHash, merkleRoot [32]byte) Header {
	return Header{
		Version:      1,
		Height:       height,
		Timestamp:    time.Now().Unix(),
		PreviousHash: previousHash,
		MerkleRoot:   merkleRoot,
	}
}
