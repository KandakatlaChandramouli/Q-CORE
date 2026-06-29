package blockchain

import (
	"sync"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/block"
)

type Blockchain struct {
	mu sync.RWMutex

	blocks []block.Block
}

func New() *Blockchain {
	return &Blockchain{
		blocks: make([]block.Block, 0),
	}
}

func (bc *Blockchain) Height() uint64 {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	return uint64(len(bc.blocks))
}

func (bc *Blockchain) AddBlock(b block.Block) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	bc.blocks = append(bc.blocks, b)
}

func (bc *Blockchain) Latest() (block.Block, bool) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if len(bc.blocks) == 0 {
		return block.Block{}, false
	}

	return bc.blocks[len(bc.blocks)-1], true
}
