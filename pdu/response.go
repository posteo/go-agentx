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
	"time"

	"gopkg.in/errgo.v1"
)

// Response defines the pdu response packet.
type Response struct {
	UpTime    time.Duration
	Error     Error
	Index     uint16
	Variables Variables
}

// Type returns the pdu packet type.
func (r *Response) Type() Type {
	return TypeResponse
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (r *Response) MarshalBinary() ([]byte, error) {
	buffer := &bytes.Buffer{}

	upTime := uint32(r.UpTime.Seconds() / 100)
	binary.Write(buffer, binary.LittleEndian, &upTime)
	binary.Write(buffer, binary.LittleEndian, &r.Error)
	binary.Write(buffer, binary.LittleEndian, &r.Index)

	vBytes, err := r.Variables.MarshalBinary()
	if err != nil {
		return nil, errgo.Mask(err)
	}
	buffer.Write(vBytes)

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (r *Response) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)

	upTime := uint32(0)
	if err := binary.Read(buffer, binary.LittleEndian, &upTime); err != nil {
		return errgo.Mask(err)
	}
	r.UpTime = time.Second * time.Duration(upTime*100)
	if err := binary.Read(buffer, binary.LittleEndian, &r.Error); err != nil {
		return errgo.Mask(err)
	}
	if err := binary.Read(buffer, binary.LittleEndian, &r.Index); err != nil {
		return errgo.Mask(err)
	}
	if err := r.Variables.UnmarshalBinary(data[8:]); err != nil {
		return errgo.Mask(err)
	}

	return nil
}
