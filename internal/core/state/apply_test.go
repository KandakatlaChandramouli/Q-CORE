package state

import (
	"testing"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/transaction"
)

func TestApplyTransaction(t *testing.T) {
	st := New()

	tx := transaction.New()

	if err := st.ApplyTransaction(tx); err != nil {
		t.Fatal(err)
	}
}
