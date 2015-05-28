package pdu

import "encoding"

// Packet defines a general interface for a pdu packet.
type Packet interface {
	TypeOwner
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
