// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

// AllocateIndex defiens the pdu allocate index packet.
type AllocateIndex struct {
	Variables Variables
}

// Type returns the pdu packet type.
func (ai *AllocateIndex) Type() Type {
	return TypeIndexAllocate
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (ai *AllocateIndex) MarshalBinary() ([]byte, error) {
	data, err := ai.Variables.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (ai *AllocateIndex) UnmarshalBinary(data []byte) error {
	return nil
}
