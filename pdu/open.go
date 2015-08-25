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

import (
	"github.com/posteo/go-agentx/marshaler"
	"gopkg.in/errgo.v1"
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
		return nil, errgo.Mask(err)
	}

	return combinedBytes, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (o *Open) UnmarshalBinary(data []byte) error {
	return nil
}
