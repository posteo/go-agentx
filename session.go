package agentx

import (
	"bufio"

	"github.com/juju/errgo"
	"github.com/posteo/go-agentx/pdu"
)

// Session defines an agentx session.
type Session struct {
	readWriter *bufio.ReadWriter
	sessionID  uint32
}

// ID returns the session id.
func (s *Session) ID() uint32 {
	return s.sessionID
}

// Close tears down the session with the master agent.
func (s *Session) Close() error {
	request := &pdu.Close{Reason: pdu.ReasonShutdown}

	if err := WritePacket(s.readWriter, request); err != nil {
		return errgo.Mask(err)
	}
	if err := s.readWriter.Flush(); err != nil {
		return errgo.Mask(err)
	}

	_, _, err := ReadPacket(s.readWriter)
	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}
