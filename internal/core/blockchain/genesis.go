package blockchain

import (
	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/block"
)

func (bc *Blockchain) InitGenesis() {
	if bc.Height() != 0 {
		return
	}

	genesis := block.New(
		0,
		[32]byte{},
		nil,
	)

	bc.AddBlock(genesis)
}
