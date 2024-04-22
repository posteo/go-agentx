package pdu

type GetBatch struct {
	SearchRanges Ranges
}

func (g *GetBatch) Type() Type {
	return TypeGet
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (g *GetBatch) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (g *GetBatch) UnmarshalBinary(data []byte) error {
	if err := g.SearchRanges.UnmarshalBinary(data); err != nil {
		return err
	}
	return nil
}
