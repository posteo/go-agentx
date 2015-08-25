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

import (
	"fmt"
	"strings"
)

// The various pdu packet flags.
const (
	FlagInstanceRegistration Flags = 1 << 0
	FlagNewIndex             Flags = 1 << 1
	FlagAnyIndex             Flags = 1 << 2
	FlagNonDefaultContext    Flags = 1 << 3
	FlagNetworkByteOrder     Flags = 1 << 4
)

// Flags defines pdu packet flags.
type Flags byte

func (f Flags) String() string {
	result := []string{}
	if f&FlagInstanceRegistration != 0 {
		result = append(result, "FlagInstanceRegistration")
	}
	if f&FlagNewIndex != 0 {
		result = append(result, "FlagNewIndex")
	}
	if f&FlagAnyIndex != 0 {
		result = append(result, "FlagAnyIndex")
	}
	if f&FlagNonDefaultContext != 0 {
		result = append(result, "FlagNonDefaultContext")
	}
	if f&FlagNetworkByteOrder != 0 {
		result = append(result, "FlagNetworkByteOrder")
	}
	if len(result) == 0 {
		return "(FlagNone)"
	}
	return fmt.Sprintf("(%s)", strings.Join(result, " | "))
}
