package codec

import (
	"bytes"
	"encoding/gob"
)

func Encode(v any) ([]byte, error) {
	var buf bytes.Buffer

	err := gob.NewEncoder(&buf).Encode(v)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(data []byte, v any) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}
