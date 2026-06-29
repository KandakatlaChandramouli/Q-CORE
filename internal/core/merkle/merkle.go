package merkle

import "crypto/sha256"

func Root(leaves [][]byte) [32]byte {
	if len(leaves) == 0 {
		return sha256.Sum256(nil)
	}

	hashes := make([][32]byte, len(leaves))

	for i := range leaves {
		hashes[i] = sha256.Sum256(leaves[i])
	}

	for len(hashes) > 1 {
		if len(hashes)%2 == 1 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}

		next := make([][32]byte, 0, len(hashes)/2)

		for i := 0; i < len(hashes); i += 2 {
			buf := append(hashes[i][:], hashes[i+1][:]...)
			next = append(next, sha256.Sum256(buf))
		}

		hashes = next
	}

	return hashes[0]
}
