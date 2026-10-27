// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import "github.com/posteo/go-agentx/value"

// Get defines the pdu get packet. An AgentX Get carries a SearchRangeList
// (RFC 2741 6.2.5), one range per requested variable; the master agent
// sends several when a manager asks for several OIDs in one request.
type Get struct {
	SearchRanges Ranges
}

// GetOID returns the first oid.
func (g *Get) GetOID() value.OID {
	if len(g.SearchRanges) == 0 {
		return nil
	}
	return g.SearchRanges[0].From.GetIdentifier()
}

// SetOID sets the provided oid as the only requested variable.
func (g *Get) SetOID(oid value.OID) {
	r := Range{}
	r.From.SetIdentifier(oid)
	g.SearchRanges = Ranges{r}
}

// Type returns the pdu packet type.
func (g *Get) Type() Type {
	return TypeGet
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (g *Get) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (g *Get) UnmarshalBinary(data []byte) error {
	if err := g.SearchRanges.UnmarshalBinary(data); err != nil {
		return err
	}
	return nil
}
