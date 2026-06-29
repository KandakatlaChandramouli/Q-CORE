package state

import (
	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/transaction"
)

func (s *State) ApplyTransaction(tx transaction.Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Placeholder until the address migration is complete.
	// The implementation will be completed when Transaction
	// uses Address instead of PublicKey.
	_ = tx

	return nil
}
