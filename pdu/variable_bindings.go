package pdu

import "github.com/juju/errgo"

// VariableBindings defines a list of variable bindings.
type VariableBindings struct {
	Items []*VariableBinding
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (v *VariableBindings) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (v *VariableBindings) UnmarshalBinary(data []byte) error {
	v.Items = make([]*VariableBinding, 0)
	for offset := 0; offset < len(data); {
		variableBinding := &VariableBinding{}
		if err := variableBinding.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		v.Items = append(v.Items, variableBinding)
		offset += variableBinding.ByteSize()
	}
	return nil
}
