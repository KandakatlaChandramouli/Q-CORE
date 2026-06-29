package badger

import db "github.com/dgraph-io/badger/v4"

type Options struct {
	Path     string
	InMemory bool
}

func (o Options) badgerOptions() db.Options {
	if o.InMemory {
		return db.DefaultOptions("").WithInMemory(true)
	}

	opts := db.DefaultOptions(o.Path)
	opts.Dir = o.Path
	opts.ValueDir = o.Path

	return opts
}
