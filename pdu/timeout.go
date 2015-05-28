package pdu

import "time"

// Timeout defines the pdu timeout packet.
type Timeout struct {
	Duration time.Duration
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (t *Timeout) MarshalBinary() ([]byte, error) {
	return []byte{byte(t.Duration.Seconds()), 0x00, 0x00, 0x00}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (t *Timeout) UnmarshalBinary(data []byte) error {
	return nil
}

func (t Timeout) String() string {
	return t.Duration.String()
}
