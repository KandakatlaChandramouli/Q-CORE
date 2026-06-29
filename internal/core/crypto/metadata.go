package crypto

// Metadata describes the immutable characteristics of a signature algorithm.
type Metadata struct {
	Algorithm         Algorithm
	PublicKeyBytes    int
	SignatureBytes    int
	SecurityBits      int
	Standard          string
	VariableSignature bool
	Deterministic     bool
}
