package badger

import (
	"context"
	"sync"

	db "github.com/dgraph-io/badger/v4"
)

type Engine struct {
	mu   sync.RWMutex
	opts Options
	db   *db.DB
}

func New(opts Options) *Engine {
	return &Engine{
		opts: opts,
	}
}

func (e *Engine) Open(context.Context) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.db != nil {
		return nil
	}

	database, err := db.Open(e.opts.badgerOptions())
	if err != nil {
		return err
	}

	e.db = database

	return nil
}

func (e *Engine) Close() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.db == nil {
		return nil
	}

	err := e.db.Close()
	e.db = nil

	return err
}
