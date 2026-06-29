package codec

import "testing"

type sample struct {
	A int
	B string
}

func TestCodec(t *testing.T) {
	in := sample{
		A: 42,
		B: "QCORE",
	}

	data, err := Encode(in)
	if err != nil {
		t.Fatal(err)
	}

	var out sample

	if err := Decode(data, &out); err != nil {
		t.Fatal(err)
	}

	if out != in {
		t.Fatal("codec mismatch")
	}
}
