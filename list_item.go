package agentx

import "github.com/posteo/go-agentx/pdu"

// ListItem defines an item of the list handler.
type ListItem struct {
	Type  pdu.VariableType
	Value interface{}
}
