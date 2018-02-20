// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import "github.com/posteo/go-agentx/value"

// Get defines the pdu get packet.
type Get struct {
	SearchRange Range
}

// GetOID returns the oid.
func (g *Get) GetOID() value.OID {
	return g.SearchRange.From.GetIdentifier()
}

// SetOID sets the provided oid.
func (g *Get) SetOID(oid value.OID) {
	g.SearchRange.From.SetIdentifier(oid)
}

// Type returns the pdu packet type.
func (g *Get) Type() Type {
	return TypeGet
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (g *Get) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (g *Get) UnmarshalBinary(data []byte) error {
	if err := g.SearchRange.UnmarshalBinary(data); err != nil {
		return err
	}
	return nil
}
