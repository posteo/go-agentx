// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package value_test

import (
	"testing"

	. "github.com/posteo/go-agentx/test"
	. "github.com/posteo/go-agentx/value"
)

func TestCommonPrefix(t *testing.T) {
	oid := MustParseOID("1.3.6.1.2")
	result := oid.CommonPrefix(MustParseOID("1.3.6.1.4"))
	AssertEquals(t, MustParseOID("1.3.6.1"), result)
}

func TestCompareOIDs_Less(t *testing.T) {
	oid1 := OID{1, 3, 6, 1, 2}
	oid2 := OID{1, 3, 6, 1, 4}

	// oid1 < oid2
	expected := -1
	AssertEquals(t, expected, CompareOIDs(oid1, oid2))
}

func TestCompareOIDs_Greater(t *testing.T) {
	oid1 := OID{1, 3, 6, 1, 2}
	oid2 := OID{1, 3, 6, 1, 4}

	// oid2 > oid1
	expected := 1
	AssertEquals(t, expected, CompareOIDs(oid2, oid1))
}

func TestCompareOIDs_Equals(t *testing.T) {
	oid1 := OID{1, 3, 6, 1, 4}
	oid2 := OID{1, 3, 6, 1, 4}

	// oid1 == oid2
	expected := 0
	AssertEquals(t, expected, CompareOIDs(oid1, oid2))
}

func TestCompareOIDs_NilValue(t *testing.T) {
	oid1 := OID{1, 3, 6, 1, 4}
	var oid2 OID

	// oid2 is nil, thus oid1 is greater
	expected := 1
	AssertEquals(t, expected, CompareOIDs(oid1, oid2))
}

func TestSortOIDs(t *testing.T) {
	var oidList []OID
	oid1 := OID{1, 3, 6, 1}
	oid2 := OID{1, 3, 6, 5, 7}
	oid3 := OID{1, 3, 6, 1, 12}
	oid4 := OID{1, 3, 6, 5}

	oidList = append(oidList, oid1, oid2, oid3, oid4)
	SortOIDs(oidList)

	var expect []OID
	expect = append(expect, oid1, oid3, oid4, oid2)
	AssertEquals(t, expect, oidList)
}
