package agentx

import (
	"errors"
	"log"
	"time"

	"github.com/posteo/go-agentx/pdu"
	"gopkg.in/errgo.v1"
)

// Session defines an agentx session.
type Session struct {
	GetHandler     func(string) (*Item, error)
	GetNextHandler func(string, string) (*Item, error)

	client    *Client
	sessionID uint32
	timeout   time.Duration
}

// ID returns the session id.
func (s *Session) ID() uint32 {
	return s.sessionID
}

// Register registers the client under the provided rootID with the provided priority
// on the master agent.
func (s *Session) Register(priority byte, rootOID string) error {
	requestPacket := &pdu.Register{}
	requestPacket.Timeout.Duration = s.timeout
	requestPacket.Timeout.Priority = priority
	requestPacket.Subtree.SetIdentifier(rootOID)
	request := &pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket}

	response := s.request(request)
	if err := checkError(response); err != nil {
		return errgo.Mask(err)
	}
	return nil
}

// AllocateIndex allocates an index at the provided oid on the master agent.
func (s *Session) AllocateIndex(oid string) error {
	requestPacket := &pdu.AllocateIndex{}
	requestPacket.Variables.Add(pdu.VariableTypeOctetString, oid, "index")
	request := &pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket}

	response := s.request(request)
	if err := checkError(response); err != nil {
		return errgo.Mask(err)
	}
	return nil
}

// Close tears down the session with the master agent.
func (s *Session) Close() error {
	requestPacket := &pdu.Close{Reason: pdu.ReasonShutdown}

	response := s.request(&pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket})
	if err := checkError(response); err != nil {
		return errgo.Mask(err)
	}
	return nil
}

func (s *Session) open(nameOID, name string) error {
	requestPacket := &pdu.Open{}
	requestPacket.Timeout.Duration = s.timeout
	requestPacket.ID.SetIdentifier(nameOID)
	requestPacket.Description.Text = name

	response := s.request(&pdu.HeaderPacket{Header: &pdu.Header{}, Packet: requestPacket})
	if err := checkError(response); err != nil {
		return errgo.Mask(err)
	}
	s.sessionID = response.Header.SessionID

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
		if s.GetHandler == nil {
			log.Printf("warning: no get handler for session specified")
			responsePacket.Variables.Add(pdu.VariableTypeNull, requestPacket.GetOID(), nil)
		} else {
			item, err := s.GetHandler(requestPacket.GetOID())
			if err != nil {
				log.Printf("error while handling packet: %s", errgo.Details(err))
				responsePacket.Error = pdu.ErrorProcessing
			}
			if item == nil {
				responsePacket.Variables.Add(pdu.VariableTypeNoSuchObject, requestPacket.GetOID(), nil)
			} else {
				oid := item.OID
				if oid == "" {
					oid = requestPacket.GetOID()
				}
				responsePacket.Variables.Add(item.Type, oid, item.Value)
			}
		}
	case *pdu.GetNext:
		if s.GetNextHandler == nil {
			log.Printf("warning: no get next handler for session specified")
		} else {
			for _, sr := range requestPacket.SearchRanges {
				item, err := s.GetNextHandler(sr.From.GetIdentifier(), sr.To.GetIdentifier())
				if err != nil {
					log.Printf("error while handling packet: %s", errgo.Details(err))
					responsePacket.Error = pdu.ErrorProcessing
				}

				if item == nil {
					responsePacket.Variables.Add(pdu.VariableTypeEndOfMIBView, sr.From.GetIdentifier(), nil)
				} else {
					responsePacket.Variables.Add(item.Type, item.OID, item.Value)
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
