package pdu

import (
	"fmt"

	"github.com/juju/errgo"
)

// Range defines the pdu search range packet.
type Range struct {
	From ObjectIdentifier
	To   ObjectIdentifier
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
