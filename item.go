package agentx

import "github.com/posteo/go-agentx/pdu"

// Item defines a structure that holds information about
// the type and value of an item.
type Item struct {
	OID   string
	Type  pdu.VariableType
	Value interface{}
}
