package crypto

import "testing"

func TestAlgorithms(t *testing.T) {
	tests := []Algorithm{
		AlgorithmECDSAP256,
		AlgorithmMLDSA44,
	}

	for _, alg := range tests {
		if alg == "" {
			t.Fatal("algorithm must not be empty")
		}
	}
}
