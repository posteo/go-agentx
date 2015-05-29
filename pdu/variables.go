package pdu

import "github.com/juju/errgo"

// Variables defines a list of variable bindings.
type Variables []Variable

// Add adds the provided variable.
func (v *Variables) Add(t VariableType, oid string, value interface{}) {
	variable := Variable{}
	variable.Set(t, oid, value)
	*v = append(*v, variable)
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (v *Variables) MarshalBinary() ([]byte, error) {
	result := []byte{}
	for _, variable := range *v {
		data, err := variable.MarshalBinary()
		if err != nil {
			return nil, errgo.Mask(err)
		}
		result = append(result, data...)
	}
	return result, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (v *Variables) UnmarshalBinary(data []byte) error {
	*v = make([]Variable, 0)
	for offset := 0; offset < len(data); {
		variable := Variable{}
		if err := variable.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		*v = append(*v, variable)
		offset += variable.ByteSize()
	}
	return nil
}
