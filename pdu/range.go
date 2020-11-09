// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import (
	"fmt"
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
		return err
	}
	if err := r.To.UnmarshalBinary(data[r.From.ByteSize():]); err != nil {
		return err
	}
	return nil
}

func (r Range) String() string {
	result := ""
	if r.From.GetInclude() {
		result += "["
	} else {
		result += "("
	}
	result += fmt.Sprintf("%v, %v", r.From, r.To)
	if r.To.GetInclude() {
		result += "]"
	} else {
		result += ")"
	}
	return result
}
