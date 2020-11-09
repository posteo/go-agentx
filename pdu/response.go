// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import (
	"bytes"
	"encoding/binary"
	"time"
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
		return nil, err
	}
	buffer.Write(vBytes)

	return buffer.Bytes(), nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (r *Response) UnmarshalBinary(data []byte) error {
	buffer := bytes.NewBuffer(data)

	upTime := uint32(0)
	if err := binary.Read(buffer, binary.LittleEndian, &upTime); err != nil {
		return err
	}
	r.UpTime = time.Second * time.Duration(upTime*100)
	if err := binary.Read(buffer, binary.LittleEndian, &r.Error); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &r.Index); err != nil {
		return err
	}
	if err := r.Variables.UnmarshalBinary(data[8:]); err != nil {
		return err
	}

	return nil
}

func (r *Response) String() string {
	return "(response " + r.Variables.String() + ")"
}
