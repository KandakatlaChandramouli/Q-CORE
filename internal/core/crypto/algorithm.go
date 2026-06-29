package crypto

// Algorithm identifies a digital signature algorithm supported by Q-CORE.
type Algorithm string

const (
	AlgorithmECDSAP256 Algorithm = "ECDSA-P256"
	AlgorithmMLDSA44   Algorithm = "ML-DSA-44"
)
