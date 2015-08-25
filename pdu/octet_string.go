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
	"bytes"
	"encoding/binary"

	"gopkg.in/errgo.v1"
)

// OctetString defines the pdu description packet.
type OctetString struct {
	Text string
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (o *OctetString) MarshalBinary() ([]byte, error) {
	buffer := &bytes.Buffer{}

	binary.Write(buffer, binary.LittleEndian, uint32(len(o.Text)))
	buffer.WriteString(o.Text)

	for buffer.Len()%4 > 0 {
		buffer.WriteByte(0x00)
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (o *OctetString) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)

	length := uint32(0)
	if err := binary.Read(buffer, binary.LittleEndian, &length); err != nil {
		return errgo.Mask(err)
	}

	o.Text = string(data[4 : 4+length])

	return nil
}
