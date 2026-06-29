package badger

import (
	"context"
	"fmt"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/block"
	"github.com/KandakatlaChandramouli/Q-CORE/internal/storage"
	"github.com/KandakatlaChandramouli/Q-CORE/internal/storage/badger/codec"
	"github.com/dgraph-io/badger/v4"
)

func blockKey(height uint64) []byte {
	return []byte(fmt.Sprintf("block:%020d", height))
}

func (e *Engine) StoreBlock(ctx context.Context, b *block.Block) error {
	_ = ctx

	db, err := e.database()
	if err != nil {
		return err
	}

	data, err := codec.Encode(b)
	if err != nil {
		return err
	}

	return db.Update(func(txn *badger.Txn) error {
		return txn.Set(blockKey(b.Header.Height), data)
	})
}

func (e *Engine) LoadBlock(ctx context.Context, height uint64) (*block.Block, error) {
	_ = ctx

	db, err := e.database()
	if err != nil {
		return nil, err
	}

	var blk block.Block

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(blockKey(height))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return storage.ErrBlockNotFound
			}
			return err
		}

		return item.Value(func(v []byte) error {
			return codec.Decode(v, &blk)
		})
	})

	if err != nil {
		return nil, err
	}

	return &blk, nil
}
