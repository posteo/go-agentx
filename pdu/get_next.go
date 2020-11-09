// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

// GetNext defines the pdu get next packet.
type GetNext struct {
	SearchRanges Ranges
}

// Type returns the pdu packet type.
func (g *GetNext) Type() Type {
	return TypeGetNext
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (g *GetNext) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (g *GetNext) UnmarshalBinary(data []byte) error {
	if err := g.SearchRanges.UnmarshalBinary(data); err != nil {
		return err
	}
	return nil
}
