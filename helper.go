package agentx

import (
	"encoding"
	"io"

	"github.com/juju/errgo"
	"github.com/posteo/go-agentx/pdu"
)

// ReadPacket reads a header from the provided reader.
func ReadPacket(r io.Reader) (*pdu.Header, pdu.Packet, error) {
	header := &pdu.Header{}
	if err := ReadBinary(r, header, pdu.HeaderSize); err != nil {
		return nil, nil, errgo.Mask(err)
	}

	var result pdu.Packet
	switch header.Type {
	case pdu.TypeResponse:
		result = &pdu.Response{}
	}
	if err := ReadBinary(r, result, header.PayloadLength); err != nil {
		return header, nil, errgo.Mask(err)
	}

	return header, result, nil
}

// ReadBinary reads the provided number of bytes from the provided reader
// and unmarshals the result using the provided unmarshaler.
func ReadBinary(r io.Reader, m encoding.BinaryUnmarshaler, c uint32) error {
	data := make([]byte, c)

	if _, err := r.Read(data); err != nil {
		return errgo.Mask(err)
	}

	if err := m.UnmarshalBinary(data); err != nil {
		return errgo.Mask(err)
	}

	return nil
}

// WriteBinary marshals the provided marshaler and writes the result
// to the provided writer.
func WriteBinary(w io.Writer, m encoding.BinaryMarshaler) error {
	data, err := m.MarshalBinary()
	if err != nil {
		return errgo.Mask(err)
	}

	if _, err := w.Write(data); err != nil {
		return errgo.Mask(err)
	}

	return nil
}

// WritePacket marshals the provided packet, and a header and writes the result
// to the provided writer.
func WritePacket(w io.Writer, p pdu.Packet) error {
	if err := WriteBinary(w, &pdu.Headed{Packet: p}); err != nil {
		return errgo.Mask(err)
	}
	return nil
}
