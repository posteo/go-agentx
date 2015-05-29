package pdu

import "gopkg.in/errgo.v1"

// Ranges defines the pdu search range list packet.
type Ranges []Range

// MarshalBinary returns the pdu packet as a slice of bytes.
func (r *Ranges) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (r *Ranges) UnmarshalBinary(data []byte) error {
	*r = make([]Range, 0)
	for offset := 0; offset < len(data); {
		rng := Range{}
		if err := rng.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		*r = append(*r, rng)
		offset += rng.ByteSize()
	}
	return nil
}
