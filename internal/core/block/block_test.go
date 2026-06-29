package block

import (
	"testing"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/transaction"
)

func TestNewBlock(t *testing.T) {
	tx := transaction.New()
	tx.Amount = 100

	b := New(1, [32]byte{}, []transaction.Transaction{tx})

	if len(b.Transactions) != 1 {
		t.Fatal("unexpected transaction count")
	}

	if b.Header.Height != 1 {
		t.Fatal("unexpected block height")
	}
}
