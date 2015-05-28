package pdu

import (
	"bytes"
	"encoding/binary"

	"github.com/juju/errgo"
)

// VariableBinding defines the pdu varbind packet.
type VariableBinding struct {
	Type  VariableType
	Name  ObjectIdentifier
	Value interface{}
}

// ByteSize returns the number of bytes, the binding would need in the encoded version.
func (v *VariableBinding) ByteSize() int {
	bytes, err := v.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return len(bytes)
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (v *VariableBinding) MarshalBinary() ([]byte, error) {
	buffer := &bytes.Buffer{}

	binary.Write(buffer, binary.LittleEndian, &v.Type)
	buffer.WriteByte(0x00)
	buffer.WriteByte(0x00)

	nameBytes, err := v.Name.MarshalBinary()
	if err != nil {
		return []byte{}, errgo.Mask(err)
	}
	buffer.Write(nameBytes)

	switch v.Type {
	case VariableTypeOctetString:
		octetString := &OctetString{Text: v.Value.(string)}
		octetStringBytes, err := octetString.MarshalBinary()
		if err != nil {
			return []byte{}, errgo.Mask(err)
		}
		buffer.Write(octetStringBytes)
	default:
		return []byte{}, errgo.Newf("unhandled variable type %s", v.Type)
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (v *VariableBinding) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)

	if err := binary.Read(buffer, binary.LittleEndian, &v.Type); err != nil {
		return errgo.Mask(err)
	}
	offset := 4

	if err := v.Name.UnmarshalBinary(data[offset:]); err != nil {
		return errgo.Mask(err)
	}
	offset += v.Name.ByteSize()

	switch v.Type {
	case VariableTypeOctetString:
		octetString := &OctetString{}
		if err := octetString.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		v.Value = octetString.Text
	default:
		return errgo.Newf("unhandled variable type %s", v.Type)
	}

	return nil
}
