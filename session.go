package agentx

import (
	"errors"
	"log"
	"time"

	"github.com/juju/errgo"
	"github.com/posteo/go-agentx/pdu"
)

// Session defines an agentx session.
type Session struct {
	client    *Client
	sessionID uint32
	timeout   time.Duration
}

// ID returns the session id.
func (s *Session) ID() uint32 {
	return s.sessionID
}

// Close tears down the session with the master agent.
func (s *Session) Close() error {
	requestPacket := &pdu.Close{Reason: pdu.ReasonShutdown}

	response := s.request(&pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket})
	if err := checkError(response); err != nil {
		return errgo.Mask(err)
	}
	log.Printf("close response: %v", response)

	return nil
}

func (s *Session) open(nameOID, name string) error {
	requestPacket := &pdu.Open{}
	requestPacket.Timeout.Duration = s.timeout
	requestPacket.ID.SetByOID(nameOID)
	requestPacket.Description.Text = name

	response := s.request(&pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket})
	if err := checkError(response); err != nil {
		return errgo.Mask(err)
	}
	s.sessionID = response.Header.SessionID

	return nil
}

func (s *Session) register(rootOID string) error {
	requestPacket := &pdu.Register{}
	requestPacket.Timeout.Duration = s.timeout
	requestPacket.Timeout.Priority = 127
	requestPacket.Subtree.SetByOID(rootOID)

	response := s.request(&pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket})
	if err := checkError(response); err != nil {
		return errgo.Mask(err)
	}
	log.Printf("register response: %v", response)

	return nil
}

func (s *Session) request(hp *pdu.HeaderPacket) *pdu.HeaderPacket {
	hp.Header.SessionID = s.sessionID
	return s.client.request(hp)
}

func checkError(hp *pdu.HeaderPacket) error {
	response, ok := hp.Packet.(*pdu.Response)
	if !ok {
		return nil
	}
	if response.Error == pdu.ErrorNone {
		return nil
	}
	return errors.New(response.Error.String())
}
