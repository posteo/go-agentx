// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// Session defines an agentx session.
type Session struct {
	Handler Handler

	client    *Client
	sessionID uint32
	timeout   time.Duration

	openRequestPacket     *pdu.HeaderPacket
	registerRequestPacket *pdu.HeaderPacket
}

// ID returns the session id.
func (s *Session) ID() uint32 {
	return s.sessionID
}

// Register registers the client under the provided rootID with the provided priority
// on the master agent.
func (s *Session) Register(priority byte, baseOID value.OID) error {
	if s.registerRequestPacket != nil {
		return fmt.Errorf("session is already registered")
	}

	requestPacket := &pdu.Register{}
	requestPacket.Timeout.Duration = s.timeout
	requestPacket.Timeout.Priority = priority
	requestPacket.Subtree.SetIdentifier(baseOID)
	request := &pdu.HeaderPacket{Header: &pdu.Header{Type: pdu.TypeRegister}, Packet: requestPacket}

	response := s.request(request)
	if err := checkError(response); err != nil {
		return err
	}
	s.registerRequestPacket = request
	return nil
}

// Unregister removes the registration for the provided subtree.
func (s *Session) Unregister(priority byte, baseOID value.OID) error {
	if s.registerRequestPacket == nil {
		return fmt.Errorf("session is not registered")
	}

	requestPacket := &pdu.Unregister{}
	requestPacket.Timeout.Duration = s.timeout
	requestPacket.Timeout.Priority = priority
	requestPacket.Subtree.SetIdentifier(baseOID)
	request := &pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket}

	response := s.request(request)
	if err := checkError(response); err != nil {
		return err
	}
	s.registerRequestPacket = nil
	return nil
}

// Close tears down the session with the master agent.
func (s *Session) Close() error {
	requestPacket := &pdu.Close{Reason: pdu.ReasonShutdown}

	response := s.request(&pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket})
	if err := checkError(response); err != nil {
		return err
	}
	return nil
}

func (s *Session) open(nameOID value.OID, name string) error {
	requestPacket := &pdu.Open{}
	requestPacket.Timeout.Duration = s.timeout
	requestPacket.ID.SetIdentifier(nameOID)
	requestPacket.Description.Text = name
	request := &pdu.HeaderPacket{Header: &pdu.Header{Type: pdu.TypeOpen}, Packet: requestPacket}

	response := s.request(request)
	if err := checkError(response); err != nil {
		return err
	}
	s.sessionID = response.Header.SessionID
	s.openRequestPacket = request
	return nil
}

func (s *Session) reopen() error {
	if s.openRequestPacket != nil {
		response := s.request(s.openRequestPacket)
		if err := checkError(response); err != nil {
			return err
		}
		s.sessionID = response.Header.SessionID
	}

	if s.registerRequestPacket != nil {
		response := s.request(s.registerRequestPacket)
		if err := checkError(response); err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) request(hp *pdu.HeaderPacket) *pdu.HeaderPacket {
	hp.Header.SessionID = s.sessionID
	return s.client.request(hp)
}

func (s *Session) handle(request *pdu.HeaderPacket) *pdu.HeaderPacket {
	responseHeader := &pdu.Header{}
	responseHeader.SessionID = request.Header.SessionID
	responseHeader.TransactionID = request.Header.TransactionID
	responseHeader.PacketID = request.Header.PacketID
	responsePacket := &pdu.Response{}

	switch requestPacket := request.Packet.(type) {
	case *pdu.Get:
		if s.Handler == nil {
			log.Printf("warning: no handler for session specified")
			responsePacket.Variables.Add(requestPacket.GetOID(), pdu.VariableTypeNull, nil)
		} else {
			oid, t, v, err := s.Handler.Get(requestPacket.GetOID())
			if err != nil {
				log.Printf("error while handling packet: %v", err)
				responsePacket.Error = pdu.ErrorProcessing
			}
			if oid == nil {
				responsePacket.Variables.Add(requestPacket.GetOID(), pdu.VariableTypeNoSuchObject, nil)
			} else {
				responsePacket.Variables.Add(oid, t, v)
			}
		}
	case *pdu.GetNext:
		if s.Handler == nil {
			log.Printf("warning: no handler for session specified")
		} else {
			for _, sr := range requestPacket.SearchRanges {
				oid, t, v, err := s.Handler.GetNext(sr.From.GetIdentifier(), (sr.From.Include == 1), sr.To.GetIdentifier())
				if err != nil {
					log.Printf("error while handling packet: %v", err)
					responsePacket.Error = pdu.ErrorProcessing
				}

				if oid == nil {
					responsePacket.Variables.Add(sr.From.GetIdentifier(), pdu.VariableTypeEndOfMIBView, nil)
				} else {
					responsePacket.Variables.Add(oid, t, v)
				}
			}
		}
	default:
		log.Printf("cannot handle unrequested packet: %v", request)
		responsePacket.Error = pdu.ErrorProcessing
	}

	return &pdu.HeaderPacket{Header: responseHeader, Packet: responsePacket}
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
