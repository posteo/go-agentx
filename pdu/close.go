// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

// Close defines the pdu close packet.
type Close struct {
	Reason Reason
}

// Type returns the pdu packet type.
func (c *Close) Type() Type {
	return TypeClose
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (c *Close) MarshalBinary() ([]byte, error) {
	return []byte{byte(c.Reason), 0x00, 0x00, 0x00}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (c *Close) UnmarshalBinary(data []byte) error {
	c.Reason = Reason(data[0])
	return nil
}
