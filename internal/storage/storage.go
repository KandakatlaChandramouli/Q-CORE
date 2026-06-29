package storage

import (
	"context"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/account"
	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/address"
	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/block"
)

type Engine interface {
	Open(ctx context.Context) error
	Close() error

	StoreBlock(ctx context.Context, b *block.Block) error
	LoadBlock(ctx context.Context, height uint64) (*block.Block, error)

	StoreAccount(ctx context.Context, addr address.Address, acc *account.Account) error
	LoadAccount(ctx context.Context, addr address.Address) (*account.Account, error)

	Height(ctx context.Context) (uint64, error)

	Sync(ctx context.Context) error
}
