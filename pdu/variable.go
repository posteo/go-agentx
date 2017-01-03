/*
go-agentx
Copyright (C) 2015 Philipp Brüll <bruell@simia.tech>

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301
USA
*/

package pdu

import (
	"bytes"
	"encoding/binary"
	"net"
	"time"

	"github.com/martinclaro/go-agentx/value"
	"gopkg.in/errgo.v1"
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
		return nil, errgo.Mask(err)
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
			return nil, errgo.Mask(err)
		}
		buffer.Write(octetStringBytes)
	case VariableTypeNull, VariableTypeNoSuchObject, VariableTypeNoSuchInstance, VariableTypeEndOfMIBView:
		break
	case VariableTypeObjectIdentifier:
		targetOID, err := value.ParseOID(v.Value.(string))
		if err != nil {
			return nil, errgo.Mask(err)
		}

		oi := &ObjectIdentifier{}
		oi.SetIdentifier(targetOID)
		oiBytes, err := oi.MarshalBinary()
		if err != nil {
			return nil, errgo.Mask(err)
		}
		buffer.Write(oiBytes)
	case VariableTypeIPAddress:
		ip := v.Value.(net.IP)
		octetString := &OctetString{Text: string(ip)}
		octetStringBytes, err := octetString.MarshalBinary()
		if err != nil {
			return nil, errgo.Mask(err)
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
			return nil, errgo.Mask(err)
		}
		buffer.Write(octetStringBytes)
	case VariableTypeCounter64:
		value := v.Value.(uint64)
		binary.Write(buffer, binary.LittleEndian, &value)
	default:
		return nil, errgo.Newf("unhandled variable type %s", v.Type)
	}

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (v *Variable) UnmarshalBinary(data []byte) error {
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
	case VariableTypeInteger:
		value := int32(0)
		if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
			return errgo.Mask(err)
		}
		v.Value = value
	case VariableTypeOctetString:
		octetString := &OctetString{}
		if err := octetString.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		v.Value = octetString.Text
	case VariableTypeNull, VariableTypeNoSuchObject, VariableTypeNoSuchInstance, VariableTypeEndOfMIBView:
		v.Value = nil
	case VariableTypeObjectIdentifier:
		oid := &ObjectIdentifier{}
		if err := oid.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		v.Value = oid.GetIdentifier()
	case VariableTypeIPAddress:
		octetString := &OctetString{}
		if err := octetString.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		v.Value = net.IP(octetString.Text)
	case VariableTypeCounter32, VariableTypeGauge32:
		value := uint32(0)
		if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
			return errgo.Mask(err)
		}
		v.Value = value
	case VariableTypeTimeTicks:
		value := uint32(0)
		if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
			return errgo.Mask(err)
		}
		v.Value = time.Duration(value) * time.Second / 100
	case VariableTypeOpaque:
		octetString := &OctetString{}
		if err := octetString.UnmarshalBinary(data[offset:]); err != nil {
			return errgo.Mask(err)
		}
		v.Value = []byte(octetString.Text)
	case VariableTypeCounter64:
		value := uint64(0)
		if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
			return errgo.Mask(err)
		}
		v.Value = value
	default:
		return errgo.Newf("unhandled variable type %s", v.Type)
	}

	return nil
}
