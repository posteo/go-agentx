package pdu

import (
	"fmt"

	"gopkg.in/errgo.v1"
)

// HeaderPacket defines a container structure for a header and a packet.
type HeaderPacket struct {
	Header *Header
	Packet Packet
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (hp *HeaderPacket) MarshalBinary() ([]byte, error) {
	payloadBytes, err := hp.Packet.MarshalBinary()
	if err != nil {
		return nil, errgo.Mask(err)
	}

	hp.Header.Version = 1
	hp.Header.Type = hp.Packet.Type()
	hp.Header.PayloadLength = uint32(len(payloadBytes))

	result, err := hp.Header.MarshalBinary()
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return append(result, payloadBytes...), nil
}

func (hp *HeaderPacket) String() string {
	return fmt.Sprintf("[HEAD %v BODY %v]", hp.Header, hp.Packet)
}
