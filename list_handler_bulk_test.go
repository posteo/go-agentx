// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

func testOIDs(n int) []string {
	r := rand.New(rand.NewSource(1))
	var out []string
	for i := 0; i < n; i++ {
		out = append(out, fmt.Sprintf("1.3.6.1.2.1.4.%d.1.%d.%d.%d", 20+r.Intn(3), 1+r.Intn(5), 1000+r.Intn(12), r.Intn(300)))
	}
	return out
}

// linearGetNext is the previous GetNext: the reference for the binary search.
func linearGetNext(l *ListHandler, from value.OID, includeFrom bool, to value.OID) value.OID {
	for _, oid := range l.oids {
		if oidWithin(oid, from, includeFrom, to) {
			return oid
		}
	}
	return nil
}

func TestAddUnsortedMatchesAdd(t *testing.T) {
	oids := testOIDs(2000)
	a, b := &ListHandler{}, &ListHandler{}
	for i, o := range oids {
		a.Add(o).Value = int32(i)
		b.AddUnsorted(o).Value = int32(i)
	}
	b.Sort()
	if len(a.oids) != len(b.oids) {
		t.Fatalf("len %d vs %d", len(a.oids), len(b.oids))
	}
	for i := range a.oids {
		if value.CompareOIDs(a.oids[i], b.oids[i]) != 0 {
			t.Fatalf("oid %d: %s vs %s", i, a.oids[i], b.oids[i])
		}
	}
}

func TestGetNextBinaryMatchesLinear(t *testing.T) {
	l := &ListHandler{}
	for _, o := range testOIDs(3000) {
		it := l.AddUnsorted(o)
		it.Type = pdu.VariableTypeInteger
		it.Value = int32(1)
	}
	l.Sort()
	ctx := context.Background()
	probes := append(testOIDs(500), "1.3.6.1.2.1.4", "1.3.6.1.2.1.4.21", "1.3.6.1.2.1.4.22.1.5", "1.3.6.1.2.1.5", "0")
	tos := []value.OID{nil, value.MustParseOID("1.3.6.1.2.1.4.21"), value.MustParseOID("1.3.6.1.2.1.5"), value.MustParseOID("1.3.6.1.2.1.4.20.1.3")}
	for _, p := range probes {
		from := value.MustParseOID(p)
		for _, inc := range []bool{false, true} {
			for _, to := range tos {
				want := linearGetNext(l, from, inc, to)
				got, typ, _, _ := l.GetNext(ctx, from, inc, to)
				if want == nil {
					if typ != pdu.VariableTypeNoSuchObject {
						t.Errorf("from %s inc %t to %s: got %s, want none", from, inc, to, got)
					}
					continue
				}
				if value.CompareOIDs(got, want) != 0 {
					t.Errorf("from %s inc %t to %s: got %s, want %s", from, inc, to, got, want)
				}
			}
		}
	}
}

// TestBulkHandlerMatchesListHandler checks BulkHandler against ListHandler
// on the same rows, duplicates included.
func TestBulkHandlerMatchesListHandler(t *testing.T) {
	oids := testOIDs(3000)
	l, b := &ListHandler{}, &BulkHandler{}
	for i, o := range oids {
		it := l.Add(o)
		it.Type = pdu.VariableTypeInteger
		it.Value = int32(i)
		b.Add(value.MustParseOID(o), pdu.VariableTypeInteger, int32(i))
	}
	b.Sort()
	ctx := context.Background()
	for _, o := range append(testOIDs(500), oids[:200]...) {
		oid := value.MustParseOID(o)
		_, lt, lv, _ := l.Get(ctx, oid)
		_, bt, bv, _ := b.Get(ctx, oid)
		if lt != bt || lv != bv {
			t.Errorf("Get %s: list %v %v, bulk %v %v", o, lt, lv, bt, bv)
		}
		for _, inc := range []bool{false, true} {
			for _, to := range []value.OID{nil, value.MustParseOID("1.3.6.1.2.1.4.21"), value.MustParseOID("1.3.6.1.2.1.5")} {
				lo, lt, lv, _ := l.GetNext(ctx, oid, inc, to)
				bo, bt, bv, _ := b.GetNext(ctx, oid, inc, to)
				if lo.String() != bo.String() || lt != bt || lv != bv {
					t.Errorf("GetNext %s inc %t to %s: list %s %v, bulk %s %v", o, inc, to, lo, lv, bo, bv)
				}
			}
		}
	}
	empty := &BulkHandler{}
	if oid, typ, _, _ := empty.GetNext(ctx, value.MustParseOID("1"), false, value.MustParseOID("2")); oid != nil || typ != pdu.VariableTypeNoSuchObject {
		t.Errorf("empty handler answered %s", oid)
	}
}
