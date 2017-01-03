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
	"github.com/martinclaro/go-agentx/value"
	"gopkg.in/errgo.v1"
)

// Variables defines a list of variable bindings.
type Variables []Variable

// Add adds the provided variable.
func (v *Variables) Add(oid value.OID, t VariableType, value interface{}) {
	variable := Variable{}
	variable.Set(oid, t, value)
	*v = append(*v, variable)
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (v *Variables) MarshalBinary() ([]byte, error) {
	result := []byte{}
	for _, variable := range *v {
		data, err := variable.MarshalBinary()
		if err != nil {
			return nil, errgo.Mask(err)
		}
		result = append(result, data...)
	}
	return result, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (v *Variables) UnmarshalBinary(data []byte) error {
	*v = make([]Variable, 0)
	for offset := 0; offset < len(data); {
		variable := Variable{}
		if err := variable.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		*v = append(*v, variable)
		offset += variable.ByteSize()
	}
	return nil
}
