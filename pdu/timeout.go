// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import "time"

// Timeout defines the pdu timeout packet.
type Timeout struct {
	Duration time.Duration
	Priority byte
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (t *Timeout) MarshalBinary() ([]byte, error) {
	return []byte{byte(t.Duration.Seconds()), t.Priority, 0x00, 0x00}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (t *Timeout) UnmarshalBinary(data []byte) error {
	t.Duration = time.Duration(data[0]) * time.Second
	t.Priority = data[1]
	return nil
}

func (t Timeout) String() string {
	return t.Duration.String()
}
