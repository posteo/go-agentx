/*
go-agentx
Copyright (C) 2015 Philipp Brüll <bruell@simia.tech>

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301
USA
*/

package pdu

import "github.com/martinclaro/go-agentx/value"

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
