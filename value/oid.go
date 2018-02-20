// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package value

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/errgo.v1"
)

// OID defines an OID.
type OID []uint32

// ParseOID parses the provided string and returns a valid oid. If one of the
// subidentifers canot be parsed to an uint32, the function will panic.
func ParseOID(text string) (OID, error) {
	var result OID

	parts := strings.Split(text, ".")
	for _, part := range parts {
		subidentifier, err := strconv.ParseUint(part, 10, 32)
		if err != nil {
			return nil, errgo.Notef(err, "could not subidentifier [%s] to uint32", part)
		}
		result = append(result, uint32(subidentifier))
	}

	return result, nil
}

// MustParseOID works like ParseOID expect it panics on a parsing error.
func MustParseOID(text string) OID {
	result, err := ParseOID(text)
	if err != nil {
		panic(err)
	}
	return result
}

// First returns the first n subidentifiers as a new oid.
func (o OID) First(count int) OID {
	return o[:count]
}

// CommonPrefix compares the oid with the provided one and
// returns a new oid containing all matching prefix subidentifiers.
func (o OID) CommonPrefix(other OID) OID {
	matchCount := 0

	for index, subidentifier := range o {
		if index >= len(other) || subidentifier != other[index] {
			break
		}
		matchCount++
	}

	return o[:matchCount]
}

func (o OID) String() string {
	var parts []string

	for _, subidentifier := range o {
		parts = append(parts, fmt.Sprintf("%d", subidentifier))
	}

	return strings.Join(parts, ".")
}
