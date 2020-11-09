// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import (
	"bytes"
	"encoding/binary"

	"github.com/posteo/go-agentx/value"
)

// ObjectIdentifier defines the pdu object identifier packet.
type ObjectIdentifier struct {
	Prefix         uint8
	Include        byte
	Subidentifiers []uint32
}

// SetInclude sets the include field.
func (o *ObjectIdentifier) SetInclude(value bool) {
	if value {
		o.Include = 0x01
	} else {
		o.Include = 0x00
	}
}

// GetInclude returns true if the include field ist set, false otherwise.
func (o *ObjectIdentifier) GetInclude() bool {
	if o.Include == 0x00 {
		return false
	}
	return true
}

// SetIdentifier set the subidentifiers by the provided oid string.
func (o *ObjectIdentifier) SetIdentifier(oid value.OID) {
	o.Subidentifiers = make([]uint32, 0)

	if len(oid) > 4 && oid[0] == 1 && oid[1] == 3 && oid[2] == 6 && oid[3] == 1 {
		o.Subidentifiers = append(o.Subidentifiers, uint32(1), uint32(3), uint32(6), uint32(1), uint32(oid[4]))
		oid = oid[5:]
	}

	o.Subidentifiers = append(o.Subidentifiers, oid...)
}

// GetIdentifier returns the identifier as an oid string.
func (o *ObjectIdentifier) GetIdentifier() value.OID {
	var oid value.OID
	if o.Prefix != 0 {
		oid = append(oid, 1, 3, 6, 1, uint32(o.Prefix))
	}
	return append(oid, o.Subidentifiers...)
}

// ByteSize returns the number of bytes, the binding would need in the encoded version.
func (o *ObjectIdentifier) ByteSize() int {
	return 4 + len(o.Subidentifiers)*4
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (o *ObjectIdentifier) MarshalBinary() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{byte(len(o.Subidentifiers)), o.Prefix, o.Include, 0x00})

	for _, subidentifier := range o.Subidentifiers {
		binary.Write(buffer, binary.LittleEndian, &subidentifier)
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (o *ObjectIdentifier) UnmarshalBinary(data []byte) error {
	count := data[0]
	o.Prefix = data[1]
	o.Include = data[2]

	o.Subidentifiers = make([]uint32, 0)
	buffer := bytes.NewBuffer(data[4:])
	for index := byte(0); index < count; index++ {
		var subidentifier uint32
		if err := binary.Read(buffer, binary.LittleEndian, &subidentifier); err != nil {
			return err
		}
		o.Subidentifiers = append(o.Subidentifiers, subidentifier)
	}

	return nil
}

func (o ObjectIdentifier) String() string {
	return o.GetIdentifier().String()
}
