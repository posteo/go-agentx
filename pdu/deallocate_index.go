package pdu

// DeallocateIndex defiens the pdu deallocate index packet.
type DeallocateIndex struct {
	Variables Variables
}

// Type returns the pdu packet type.
func (di *DeallocateIndex) Type() Type {
	return TypeIndexDeallocate
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (di *DeallocateIndex) MarshalBinary() ([]byte, error) {
	data, err := di.Variables.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (di *DeallocateIndex) UnmarshalBinary(data []byte) error {
	return nil
}
