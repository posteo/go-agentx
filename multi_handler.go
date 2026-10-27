// Converged: answer AgentX requests from all sessions of a client.

package agentx

import (
	"context"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// multiHandler asks every handler. Get returns the first handler that
// knows the OID; GetNext returns the lowest OID any handler has within
// the search range.
type multiHandler []Handler

func (m multiHandler) Get(ctx context.Context, oid value.OID) (value.OID, pdu.VariableType, any, error) {
	for _, h := range m {
		roid, t, v, err := h.Get(ctx, oid)
		if err != nil {
			return roid, t, v, err
		}
		if roid != nil {
			return roid, t, v, nil
		}
	}
	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}

func (m multiHandler) GetNext(ctx context.Context, from value.OID, includeFrom bool, to value.OID) (value.OID, pdu.VariableType, any, error) {
	var (
		best  value.OID
		bestT pdu.VariableType
		bestV any
	)
	for _, h := range m {
		roid, t, v, err := h.GetNext(ctx, from, includeFrom, to)
		if err != nil {
			return roid, t, v, err
		}
		if roid != nil && (best == nil || value.CompareOIDs(roid, best) < 0) {
			best, bestT, bestV = roid, t, v
		}
	}
	if best == nil {
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}
	return best, bestT, bestV, nil
}
