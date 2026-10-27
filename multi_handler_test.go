// Converged: tests for answering from all sessions of a client.

package agentx

import (
	"context"
	"testing"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

func listWith(oids ...string) *ListHandler {
	l := &ListHandler{}
	for _, o := range oids {
		it := l.Add(o)
		it.Type = pdu.VariableTypeOctetString
		it.Value = o
	}
	return l
}

func TestMultiHandlerRoutesByOID(t *testing.T) {
	ifmib := listWith("1.3.6.1.2.1.31.1.1.1.1.1000", "1.3.6.1.2.1.31.1.1.1.1.1001")
	ent := listWith("1.3.6.1.2.1.99.1.1.1.4.201")
	// The DOM handler comes first, as when snmpd addresses an IF-MIB
	// variable to the DOM session.
	m := multiHandler{ent, ifmib}
	ctx := context.Background()

	oid, _, v, err := m.Get(ctx, value.MustParseOID("1.3.6.1.2.1.31.1.1.1.1.1001"))
	if err != nil || oid == nil || v != "1.3.6.1.2.1.31.1.1.1.1.1001" {
		t.Fatalf("Get ifName.1001 = %v %v %v", oid, v, err)
	}
	if oid, _, _, _ := m.Get(ctx, value.MustParseOID("1.3.6.1.2.1.31.1.1.1.1.9999")); oid != nil {
		t.Fatalf("Get of a missing OID returned %v", oid)
	}

	from := value.MustParseOID("1.3.6.1.2.1.31.1.1.1.1")
	to := value.MustParseOID("1.3.6.1.2.1.31.1.1.1.2")
	oid, _, _, _ = m.GetNext(ctx, from, false, to)
	if oid.String() != "1.3.6.1.2.1.31.1.1.1.1.1000" {
		t.Fatalf("GetNext = %v, want ifName.1000", oid)
	}
	// Range covering both MIBs: the lowest OID wins.
	oid, _, _, _ = m.GetNext(ctx, value.MustParseOID("1.3.6.1.2.1.31"), false, value.MustParseOID("1.3.6.1.2.1.100"))
	if oid.String() != "1.3.6.1.2.1.31.1.1.1.1.1000" {
		t.Fatalf("GetNext across MIBs = %v, want ifName.1000", oid)
	}
	oid, _, _, _ = m.GetNext(ctx, value.MustParseOID("1.3.6.1.2.1.31.1.1.1.1.1001"), false, value.MustParseOID("1.3.6.1.2.1.100"))
	if oid.String() != "1.3.6.1.2.1.99.1.1.1.4.201" {
		t.Fatalf("GetNext after the last ifName = %v, want sensor 201", oid)
	}
}
