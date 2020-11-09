// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
		return fmt.Errorf("not enough bytes (%d) to unmarshal the header (%d)", len(data), HeaderSize)
	}

	h.Version, h.Type, h.Flags = data[0], Type(data[1]), Flags(data[2])

	buffer := bytes.NewBuffer(data[4:])

	binary.Read(buffer, binary.LittleEndian, &h.SessionID)
	binary.Read(buffer, binary.LittleEndian, &h.TransactionID)
	binary.Read(buffer, binary.LittleEndian, &h.PacketID)
	binary.Read(buffer, binary.LittleEndian, &h.PayloadLength)

	return nil
}

func (h *Header) String() string {
	return "(header " + h.Type.String() + ")"
}
