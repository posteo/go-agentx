// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import (
	"github.com/posteo/go-agentx/marshaler"
)

// Open defines a pdu open packet.
type Open struct {
	Timeout     Timeout
	ID          ObjectIdentifier
	Description OctetString
}

// Type returns the pdu packet type.
func (o *Open) Type() Type {
	return TypeOpen
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (o *Open) MarshalBinary() ([]byte, error) {
	combined := marshaler.NewMulti(&o.Timeout, &o.ID, &o.Description)

	combinedBytes, err := combined.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return combinedBytes, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (o *Open) UnmarshalBinary(data []byte) error {
	return nil
}
