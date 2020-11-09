// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/posteo/go-agentx/value"
)

// Variable defines the pdu varbind packet.
type Variable struct {
	Type  VariableType
	Name  ObjectIdentifier
	Value interface{}
}

// Set sets the variable.
func (v *Variable) Set(oid value.OID, t VariableType, value interface{}) {
	v.Name.SetIdentifier(oid)
	v.Type = t
	v.Value = value
}

// ByteSize returns the number of bytes, the binding would need in the encoded version.
func (v *Variable) ByteSize() int {
	bytes, err := v.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return len(bytes)
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (v *Variable) MarshalBinary() ([]byte, error) {
	buffer := &bytes.Buffer{}

	binary.Write(buffer, binary.LittleEndian, &v.Type)
	buffer.WriteByte(0x00)
	buffer.WriteByte(0x00)

	nameBytes, err := v.Name.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buffer.Write(nameBytes)

	switch v.Type {
	case VariableTypeInteger:
		value := v.Value.(int32)
		binary.Write(buffer, binary.LittleEndian, &value)
	case VariableTypeOctetString:
		octetString := &OctetString{Text: v.Value.(string)}
		octetStringBytes, err := octetString.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buffer.Write(octetStringBytes)
	case VariableTypeNull, VariableTypeNoSuchObject, VariableTypeNoSuchInstance, VariableTypeEndOfMIBView:
		break
	case VariableTypeObjectIdentifier:
		targetOID, err := value.ParseOID(v.Value.(string))
		if err != nil {
			return nil, err
		}

		oi := &ObjectIdentifier{}
		oi.SetIdentifier(targetOID)
		oiBytes, err := oi.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buffer.Write(oiBytes)
	case VariableTypeIPAddress:
		ip := v.Value.(net.IP)
		octetString := &OctetString{Text: string(ip)}
		octetStringBytes, err := octetString.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buffer.Write(octetStringBytes)
	case VariableTypeCounter32, VariableTypeGauge32:
		value := v.Value.(uint32)
		binary.Write(buffer, binary.LittleEndian, &value)
	case VariableTypeTimeTicks:
		value := uint32(v.Value.(time.Duration).Seconds() * 100)
		binary.Write(buffer, binary.LittleEndian, &value)
	case VariableTypeOpaque:
		octetString := &OctetString{Text: string(v.Value.([]byte))}
		octetStringBytes, err := octetString.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buffer.Write(octetStringBytes)
	case VariableTypeCounter64:
		value := v.Value.(uint64)
		binary.Write(buffer, binary.LittleEndian, &value)
	default:
		return nil, fmt.Errorf("unhandled variable type %s", v.Type)
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (v *Variable) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)

	if err := binary.Read(buffer, binary.LittleEndian, &v.Type); err != nil {
		return err
	}
	offset := 4

	if err := v.Name.UnmarshalBinary(data[offset:]); err != nil {
		return err
	}
	offset += v.Name.ByteSize()

	switch v.Type {
	case VariableTypeInteger:
		value := int32(0)
		if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
			return err
		}
		v.Value = value
	case VariableTypeOctetString:
		octetString := &OctetString{}
		if err := octetString.UnmarshalBinary(data[offset:]); err != nil {
			return err
		}
		v.Value = octetString.Text
	case VariableTypeNull, VariableTypeNoSuchObject, VariableTypeNoSuchInstance, VariableTypeEndOfMIBView:
		v.Value = nil
	case VariableTypeObjectIdentifier:
		oid := &ObjectIdentifier{}
		if err := oid.UnmarshalBinary(data[offset:]); err != nil {
			return err
		}
		v.Value = oid.GetIdentifier()
	case VariableTypeIPAddress:
		octetString := &OctetString{}
		if err := octetString.UnmarshalBinary(data[offset:]); err != nil {
			return err
		}
		v.Value = net.IP(octetString.Text)
	case VariableTypeCounter32, VariableTypeGauge32:
		value := uint32(0)
		if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
			return err
		}
		v.Value = value
	case VariableTypeTimeTicks:
		value := uint32(0)
		if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
			return err
		}
		v.Value = time.Duration(value) * time.Second / 100
	case VariableTypeOpaque:
		octetString := &OctetString{}
		if err := octetString.UnmarshalBinary(data[offset:]); err != nil {
			return err
		}
		v.Value = []byte(octetString.Text)
	case VariableTypeCounter64:
		value := uint64(0)
		if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
			return err
		}
		v.Value = value
	default:
		return fmt.Errorf("unhandled variable type %s", v.Type)
	}

	return nil
}

func (v *Variable) String() string {
	return fmt.Sprintf("(variable %s = %v)", v.Type, v.Value)
}
