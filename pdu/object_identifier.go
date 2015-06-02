package pdu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/errgo.v1"
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
func (o *ObjectIdentifier) SetIdentifier(value string) error {
	parts := strings.Split(value, ".")
	o.Subidentifiers = make([]uint32, 0)

	if len(parts) > 4 && parts[0] == "1" && parts[1] == "3" && parts[2] == "6" && parts[3] == "1" {
		prefix, err := strconv.Atoi(parts[4])
		if err != nil {
			return errgo.Mask(err)
		}
		o.Subidentifiers = append(o.Subidentifiers, uint32(1), uint32(3), uint32(6), uint32(1), uint32(prefix))
		parts = parts[5:]
	}

	for _, part := range parts {
		subidentifier, err := strconv.Atoi(part)
		if err != nil {
			return errgo.Mask(err)
		}
		o.Subidentifiers = append(o.Subidentifiers, uint32(subidentifier))
	}

	return nil
}

// GetIdentifier returns the identifier as an oid string.
func (o *ObjectIdentifier) GetIdentifier() string {
	var parts []string
	if o.Prefix != 0 {
		parts = append(parts, "1", "3", "6", "1", fmt.Sprintf("%d", o.Prefix))
	}
	for _, subidentifier := range o.Subidentifiers {
		parts = append(parts, fmt.Sprintf("%d", subidentifier))
	}
	return strings.Join(parts, ".")
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
			return errgo.Mask(err)
		}
		o.Subidentifiers = append(o.Subidentifiers, subidentifier)
	}

	return nil
}

func (o ObjectIdentifier) String() string {
	return o.GetIdentifier()
}
