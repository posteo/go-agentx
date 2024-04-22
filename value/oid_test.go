// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package value_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yamu-OSS/go-agentx/value"
)

func TestCommonPrefix(t *testing.T) {
	oid := value.MustParseOID("1.3.6.1.2")
	result := oid.CommonPrefix(value.MustParseOID("1.3.6.1.4"))
	assert.Equal(t, value.MustParseOID("1.3.6.1"), result)
}

func TestCompareOIDs_Less(t *testing.T) {
	oid1 := value.OID{1, 3, 6, 1, 2}
	oid2 := value.OID{1, 3, 6, 1, 4}

	// oid1 < oid2
	expected := -1
	assert.Equal(t, expected, value.CompareOIDs(oid1, oid2))
}

func TestCompareOIDs_Greater(t *testing.T) {
	oid1 := value.OID{1, 3, 6, 1, 2}
	oid2 := value.OID{1, 3, 6, 1, 4}

	// oid2 > oid1
	expected := 1
	assert.Equal(t, expected, value.CompareOIDs(oid2, oid1))
}

func TestCompareOIDs_Equals(t *testing.T) {
	oid1 := value.OID{1, 3, 6, 1, 4}
	oid2 := value.OID{1, 3, 6, 1, 4}

	// oid1 == oid2
	expected := 0
	assert.Equal(t, expected, value.CompareOIDs(oid1, oid2))
}

func TestCompareOIDs_NilValue(t *testing.T) {
	oid1 := value.OID{1, 3, 6, 1, 4}
	var oid2 value.OID

	// oid2 is nil, thus oid1 is greater
	expected := 1
	assert.Equal(t, expected, value.CompareOIDs(oid1, oid2))
}

func TestSortOIDs(t *testing.T) {
	var oidList []value.OID
	oid1 := value.OID{1, 3, 6, 1}
	oid2 := value.OID{1, 3, 6, 5, 7}
	oid3 := value.OID{1, 3, 6, 1, 12}
	oid4 := value.OID{1, 3, 6, 5}

	oidList = append(oidList, oid1, oid2, oid3, oid4)
	value.SortOIDs(oidList)

	var expect []value.OID
	expect = append(expect, oid1, oid3, oid4, oid2)
	assert.Equal(t, expect, oidList)
}
