package merkle

import "testing"

func TestRoot(t *testing.T) {
	root := Root([][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
	})

	var zero [32]byte

	if root == zero {
		t.Fatal("invalid merkle root")
	}
}
