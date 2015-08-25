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
