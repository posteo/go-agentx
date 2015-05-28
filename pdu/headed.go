package pdu

import "github.com/juju/errgo"

// Headed defines a marshaler that prepends a header to the provided packet.
type Headed struct {
	Packet Packet
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (h *Headed) MarshalBinary() ([]byte, error) {
	payloadBytes, err := h.Packet.MarshalBinary()
	if err != nil {
		return nil, errgo.Mask(err)
	}

	header := &Header{
		Version:       1,
		Type:          h.Packet.Type(),
		PayloadLength: uint32(len(payloadBytes)),
	}

	headerBytes, err := header.MarshalBinary()
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return append(headerBytes, payloadBytes...), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (h *Headed) UnmarshalBinary(data []byte) error {
	return nil
}
