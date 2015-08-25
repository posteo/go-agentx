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

import "time"

// Timeout defines the pdu timeout packet.
type Timeout struct {
	Duration time.Duration
	Priority byte
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (t *Timeout) MarshalBinary() ([]byte, error) {
	return []byte{byte(t.Duration.Seconds()), t.Priority, 0x00, 0x00}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (t *Timeout) UnmarshalBinary(data []byte) error {
	t.Duration = time.Duration(data[0]) * time.Second
	t.Priority = data[1]
	return nil
}

func (t Timeout) String() string {
	return t.Duration.String()
}
