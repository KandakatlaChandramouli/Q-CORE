package crypto

// PublicKey represents an opaque public key.
type PublicKey interface{}

// PrivateKey represents an opaque private key.
type PrivateKey interface{}

// Engine defines the behavior required from every signature algorithm.
//
// Every implementation must be deterministic with respect to its public API
// and safe for concurrent use unless documented otherwise.
type Engine interface {
	Algorithm() Algorithm

	GenerateKey() (PublicKey, PrivateKey, error)

	Sign(
		privateKey PrivateKey,
		message []byte,
	) ([]byte, error)

	Verify(
		publicKey PublicKey,
		message []byte,
		signature []byte,
	) error

	PublicKeySize() int

	SignatureSize() int
}
