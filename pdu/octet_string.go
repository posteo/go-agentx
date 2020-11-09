// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import (
	"bytes"
	"encoding/binary"
)

// OctetString defines the pdu description packet.
type OctetString struct {
	Text string
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (o *OctetString) MarshalBinary() ([]byte, error) {
	buffer := &bytes.Buffer{}

	binary.Write(buffer, binary.LittleEndian, uint32(len(o.Text)))
	buffer.WriteString(o.Text)

	for buffer.Len()%4 > 0 {
		buffer.WriteByte(0x00)
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (o *OctetString) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)

	length := uint32(0)
	if err := binary.Read(buffer, binary.LittleEndian, &length); err != nil {
		return err
	}

	o.Text = string(data[4 : 4+length])

	return nil
}
