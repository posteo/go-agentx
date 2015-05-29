package marshaler

import (
	"encoding"

	"gopkg.in/errgo.v1"
)

// Multi defines a binary marshaler that marshals all child marshalers
// and concatinate the results.
type Multi []encoding.BinaryMarshaler

// NewMulti returns a new instance of MultiBinaryMarshaler.
func NewMulti(marshalers ...encoding.BinaryMarshaler) Multi {
	return Multi(marshalers)
}

// MarshalBinary marshals all the binary marshalers and concatinates the results.
func (m Multi) MarshalBinary() ([]byte, error) {
	result := []byte{}

	for _, marshaler := range m {
		data, err := marshaler.MarshalBinary()
		if err != nil {
			return nil, errgo.Mask(err)
		}
		result = append(result, data...)
	}

	return result, nil
}
