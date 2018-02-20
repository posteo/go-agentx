// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

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
