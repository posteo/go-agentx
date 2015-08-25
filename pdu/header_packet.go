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
