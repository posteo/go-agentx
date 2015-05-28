package pdu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/juju/errgo"
)

// ObjectIdentifier defines the pdu object identifier packet.
type ObjectIdentifier struct {
	Prefix         byte
	Subidentifiers []uint32
}

// SetByOID set the subidentifiers by the provided oid string.
func (o *ObjectIdentifier) SetByOID(value string) error {
	parts := strings.Split(value, ".")
	o.Subidentifiers = make([]uint32, 0)
	for _, part := range parts {
		subidentifier, err := strconv.Atoi(part)
		if err != nil {
			return errgo.Mask(err)
		}
		o.Subidentifiers = append(o.Subidentifiers, uint32(subidentifier))
	}
	return nil
}

// ByteSize returns the number of bytes, the binding would need in the encoded version.
func (o *ObjectIdentifier) ByteSize() int {
	return 4 + len(o.Subidentifiers)*4
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (o *ObjectIdentifier) MarshalBinary() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{byte(len(o.Subidentifiers)), o.Prefix, 0x00, 0x00})

	for _, subidentifier := range o.Subidentifiers {
		if err := binary.Write(buffer, binary.LittleEndian, &subidentifier); err != nil {
			return []byte{}, errgo.Mask(err)
		}
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (o *ObjectIdentifier) UnmarshalBinary(data []byte) error {
	count := data[0]
	o.Prefix = data[1]

	o.Subidentifiers = make([]uint32, 0)
	buffer := bytes.NewBuffer(data[4:])
	for index := byte(0); index < count; index++ {
		var subidentifier uint32
		if err := binary.Read(buffer, binary.LittleEndian, &subidentifier); err != nil {
			return errgo.Mask(err)
		}
		o.Subidentifiers = append(o.Subidentifiers, subidentifier)
	}

	return nil
}

func (o ObjectIdentifier) String() string {
	result := ""
	for _, subidentifier := range o.Subidentifiers {
		result += fmt.Sprintf(".%d", subidentifier)
	}
	return fmt.Sprintf("%s (%d)", result, o.Prefix)
}
