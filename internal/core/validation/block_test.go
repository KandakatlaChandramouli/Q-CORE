package validation

import (
	"testing"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/block"
)

func TestValidateBlockHeader(t *testing.T) {
	h := block.NewHeader(1, [32]byte{}, [32]byte{})

	if err := ValidateBlockHeader(h); err != nil {
		t.Fatal(err)
	}

	h.Version = 2

	if err := ValidateBlockHeader(h); err == nil {
		t.Fatal("expected version validation failure")
	}
}
