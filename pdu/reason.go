// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

import "fmt"

// The various pdu packet reasons.
const (
	ReasonOther         Reason = 1
	ReasonParseError    Reason = 2
	ReasonProtocolError Reason = 3
	ReasonTimeouts      Reason = 4
	ReasonShutdown      Reason = 5
	ReasonByManager     Reason = 6
)

// Reason defines a reason.
type Reason byte

func (r Reason) String() string {
	switch r {
	case ReasonOther:
		return "ReasonOther"
	case ReasonParseError:
		return "ReasonParseError"
	case ReasonProtocolError:
		return "ReasonProtocolError"
	case ReasonTimeouts:
		return "ReasonTimeouts"
	case ReasonShutdown:
		return "ReasonShutdown"
	case ReasonByManager:
		return "ReasonByManager"
	}
	return fmt.Sprintf("ReasonUnknown (%d)", r)
}
