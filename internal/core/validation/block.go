package validation

import (
	"errors"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/core/block"
)

var (
	ErrInvalidVersion = errors.New("validation: invalid block version")
	ErrInvalidHeight  = errors.New("validation: invalid block height")
	ErrInvalidTime    = errors.New("validation: invalid timestamp")
)

func ValidateBlockHeader(h block.Header) error {
	if h.Version != 1 {
		return ErrInvalidVersion
	}

	if h.Timestamp <= 0 {
		return ErrInvalidTime
	}

	return nil
}
