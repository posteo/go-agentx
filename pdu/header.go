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

const (
	// HeaderSize defines the total size of a header packet.
	HeaderSize = 20
)

// Header defines a pdu packet header
type Header struct {
	Version       byte
	Type          Type
	Flags         Flags
	SessionID     uint32
	TransactionID uint32
	PacketID      uint32
	PayloadLength uint32
}

// MarshalBinary returns the pdu header as a slice of bytes.
func (h *Header) MarshalBinary() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{h.Version, byte(h.Type), byte(h.Flags), 0x00})

	binary.Write(buffer, binary.LittleEndian, h.SessionID)
	binary.Write(buffer, binary.LittleEndian, h.TransactionID)
	binary.Write(buffer, binary.LittleEndian, h.PacketID)
	binary.Write(buffer, binary.LittleEndian, h.PayloadLength)

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the header structure from the provided slice of bytes.
func (h *Header) UnmarshalBinary(data []byte) error {
	if len(data) < HeaderSize {
		return errgo.Newf("not enough bytes (%d) to unmarshal the header (%d)", len(data), HeaderSize)
	}

	h.Version, h.Type, h.Flags = data[0], Type(data[1]), Flags(data[2])

	buffer := bytes.NewBuffer(data[4:])

	binary.Read(buffer, binary.LittleEndian, &h.SessionID)
	binary.Read(buffer, binary.LittleEndian, &h.TransactionID)
	binary.Read(buffer, binary.LittleEndian, &h.PacketID)
	binary.Read(buffer, binary.LittleEndian, &h.PayloadLength)

	return nil
}
