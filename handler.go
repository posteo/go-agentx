package agentx

import (
	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// Handler defines an interface for a handler of events that
// might occure during a session.
type Handler interface {
	Get(value.OID) (value.OID, pdu.VariableType, interface{}, error)
	GetNext(value.OID, value.OID) (value.OID, pdu.VariableType, interface{}, error)
}
