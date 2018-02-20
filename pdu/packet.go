// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import "encoding"

// Packet defines a general interface for a pdu packet.
type Packet interface {
	TypeOwner
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
