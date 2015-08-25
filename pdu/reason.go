/*
go-agentx
Copyright (C) 2015 Philipp Brüll <bruell@simia.tech>

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301
USA
*/

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
