package pdu

import (
	"bytes"
	"encoding/binary"

	"github.com/juju/errgo"
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

	if err := binary.Write(buffer, binary.LittleEndian, h.SessionID); err != nil {
		return []byte{}, errgo.Mask(err)
	}
	if err := binary.Write(buffer, binary.LittleEndian, h.TransactionID); err != nil {
		return []byte{}, errgo.Mask(err)
	}
	if err := binary.Write(buffer, binary.LittleEndian, h.PacketID); err != nil {
		return []byte{}, errgo.Mask(err)
	}
	if err := binary.Write(buffer, binary.LittleEndian, h.PayloadLength); err != nil {
		return []byte{}, errgo.Mask(err)
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the header structure from the provided slice of bytes.
func (h *Header) UnmarshalBinary(data []byte) error {
	if len(data) < HeaderSize {
		return errgo.Newf("not enough bytes (%d) to unmarshal the header (%d)", len(data), HeaderSize)
	}

	h.Version, h.Type, h.Flags = data[0], Type(data[1]), Flags(data[2])

	buffer := bytes.NewBuffer(data[4:])

	if err := binary.Read(buffer, binary.LittleEndian, &h.SessionID); err != nil {
		return errgo.Mask(err)
	}
	if err := binary.Read(buffer, binary.LittleEndian, &h.TransactionID); err != nil {
		return errgo.Mask(err)
	}
	if err := binary.Read(buffer, binary.LittleEndian, &h.PacketID); err != nil {
		return errgo.Mask(err)
	}
	if err := binary.Read(buffer, binary.LittleEndian, &h.PayloadLength); err != nil {
		return errgo.Mask(err)
	}

	return nil
}
