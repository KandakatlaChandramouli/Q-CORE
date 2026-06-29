package state

import (
	"errors"
	"sync"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/account"
	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/address"
)

var (
	ErrAccountNotFound   = errors.New("state: account not found")
	ErrInsufficientFunds = errors.New("state: insufficient funds")
	ErrInvalidNonce      = errors.New("state: invalid nonce")
)

type State struct {
	mu sync.RWMutex

	accounts map[address.Address]*account.Account
}

func New() *State {
	return &State{
		accounts: make(map[address.Address]*account.Account),
	}
}

func (s *State) Create(acc account.Account) {
	s.mu.Lock()
	defer s.mu.Unlock()

	a := acc
	s.accounts[a.Address] = &a
}

func (s *State) Get(addr address.Address) (*account.Account, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	acc, ok := s.accounts[addr]
	return acc, ok
}
