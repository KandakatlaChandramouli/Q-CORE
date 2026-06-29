package crypto

type PublicKey interface{}

type PrivateKey interface{}

// Engine defines the contract implemented by every signature scheme.
type Engine interface {
	Algorithm() Algorithm

	Metadata() Metadata

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
}
