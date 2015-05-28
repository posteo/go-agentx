package pdu

import (
	"bytes"
	"encoding/binary"

	"github.com/juju/errgo"
)

// OctetString defines the pdu description packet.
type OctetString struct {
	Text string
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (o *OctetString) MarshalBinary() ([]byte, error) {
	buffer := &bytes.Buffer{}

	if err := binary.Write(buffer, binary.LittleEndian, uint32(len(o.Text))); err != nil {
		return []byte{}, errgo.Mask(err)
	}

	if _, err := buffer.WriteString(o.Text); err != nil {
		return []byte{}, errgo.Mask(err)
	}

	for buffer.Len()%4 > 0 {
		if err := buffer.WriteByte(0x00); err != nil {
			return []byte{}, errgo.Mask(err)
		}
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (o *OctetString) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)

	length := uint32(0)
	if err := binary.Read(buffer, binary.LittleEndian, &length); err != nil {
		return errgo.Mask(err)
	}

	o.Text = string(data[4 : 4+length])

	return nil
}
