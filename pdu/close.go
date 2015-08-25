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

// Close defines the pdu close packet.
type Close struct {
	Reason Reason
}

// Type returns the pdu packet type.
func (c *Close) Type() Type {
	return TypeClose
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (c *Close) MarshalBinary() ([]byte, error) {
	return []byte{byte(c.Reason), 0x00, 0x00, 0x00}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (c *Close) UnmarshalBinary(data []byte) error {
	c.Reason = Reason(data[0])
	return nil
}
