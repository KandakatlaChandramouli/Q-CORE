package block

import (
	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/merkle"
	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/transaction"
)

type Block struct {
	Header Header

	Transactions []transaction.Transaction
}

func New(
	height uint64,
	previousHash [32]byte,
	txs []transaction.Transaction,
) Block {

	leaves := make([][]byte, 0, len(txs))

	for i := range txs {
		if err := txs[i].ComputeID(); err != nil {
			panic(err)
		}

		id := txs[i].ID
		leaves = append(leaves, id[:])
	}

	root := merkle.Root(leaves)

	return Block{
		Header:       NewHeader(height, previousHash, root),
		Transactions: txs,
	}
}
