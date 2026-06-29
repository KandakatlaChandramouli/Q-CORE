// Package crypto defines the cryptographic abstraction layer used by Q-CORE.
//
// The package intentionally contains no concrete cryptographic implementation.
// Instead it defines the contracts that algorithms such as ECDSA, ML-DSA,
// and future post-quantum schemes must satisfy.
//
// This separation allows benchmarking different signature schemes while keeping
// the blockchain pipeline unchanged.
package crypto
