// Converged: regression test for multi-variable AgentX Get (RFC 2741 6.2.5).

package pdu

import (
	"testing"

	"github.com/posteo/go-agentx/value"
)

func TestGetUnmarshalsEverySearchRange(t *testing.T) {
	oids := []string{"1.3.6.1.2.1.2.2.1.2.1001", "1.3.6.1.2.1.2.2.1.10.1001", "1.3.6.1.2.1.31.1.1.1.1.1002"}
	// Ranges.MarshalBinary is a stub, so encode the SearchRangeList by
	// hand: start OID, then an empty end OID, per requested variable.
	var data []byte
	for _, s := range oids {
		from := ObjectIdentifier{}
		from.SetIdentifier(value.MustParseOID(s))
		b, err := from.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		to := ObjectIdentifier{}
		e, err := to.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		data = append(append(data, b...), e...)
	}
	g := &Get{}
	if err := g.UnmarshalBinary(data); err != nil {
		t.Fatal(err)
	}
	if len(g.SearchRanges) != len(oids) {
		t.Fatalf("got %d search ranges, want %d", len(g.SearchRanges), len(oids))
	}
	for i, s := range oids {
		if got := g.SearchRanges[i].From.GetIdentifier().String(); got != s {
			t.Errorf("range %d = %s, want %s", i, got, s)
		}
	}
	if g.GetOID().String() != oids[0] {
		t.Errorf("GetOID = %s, want %s", g.GetOID(), oids[0])
	}
}
