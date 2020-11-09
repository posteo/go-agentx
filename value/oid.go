// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package value

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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
			return nil, fmt.Errorf("parse uint [%s]: %w", part, err)
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

// CompareOIDs returns an integer comparing two OIDs lexicographically.
// The result will be 0 if oid1 == oid2, -1 if oid1 < oid2, +1 if oid1 > oid2.
func CompareOIDs(oid1, oid2 OID) int {
	if oid2 != nil {
		oid1Length := len(oid1)
		oid2Length := len(oid2)
		for i := 0; i < oid1Length && i < oid2Length; i++ {
			if oid1[i] < oid2[i] {
				return -1
			}
			if oid1[i] > oid2[i] {
				return 1
			}
		}
		if oid1Length == oid2Length {
			return 0
		} else if oid1Length < oid2Length {
			return -1
		} else {
			return 1
		}
	}
	return 1
}

// SortOIDs performs sorting of the OID list.
func SortOIDs(oids []OID) {
	sort.Slice(oids, func(i, j int) bool {
		return CompareOIDs(oids[i], oids[j]) == -1
	})
}

func (o OID) String() string {
	var parts []string

	for _, subidentifier := range o {
		parts = append(parts, fmt.Sprintf("%d", subidentifier))
	}

	return strings.Join(parts, ".")
}
