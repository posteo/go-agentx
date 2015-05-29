package pdu

import (
	"fmt"

	"gopkg.in/errgo.v1"
)

// Range defines the pdu search range packet.
type Range struct {
	From ObjectIdentifier
	To   ObjectIdentifier
}

// ByteSize returns the number of bytes, the binding would need in the encoded version.
func (r *Range) ByteSize() int {
	return r.From.ByteSize() + r.To.ByteSize()
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (r *Range) MarshalBinary() ([]byte, error) {
	r.To.SetInclude(false)
	return []byte{}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (r *Range) UnmarshalBinary(data []byte) error {
	if err := r.From.UnmarshalBinary(data); err != nil {
		return errgo.Mask(err)
	}
	if err := r.To.UnmarshalBinary(data[r.From.ByteSize():]); err != nil {
		return errgo.Mask(err)
	}
	return nil
}

func (r Range) String() string {
	return fmt.Sprintf("RANGE FROM %v TO %v", r.From, r.To)
}
