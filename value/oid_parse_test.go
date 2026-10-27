// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package value

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// splitParseOID is the previous ParseOID: the reference for the one-pass parser.
func splitParseOID(text string) (OID, error) {
	var result OID
	for _, part := range strings.Split(text, ".") {
		subidentifier, err := strconv.ParseUint(part, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("parse uint [%s]: %w", part, err)
		}
		result = append(result, uint32(subidentifier))
	}
	return result, nil
}

func TestParseOIDMatchesSplit(t *testing.T) {
	for _, in := range []string{
		"1", "0", "1.3.6.1.2.1.4.35.1.4.1007.1.4.198.18.11.4", "4294967295.1", "4294967296",
		"", ".", ".1.3", "1.3.", "1..3", "1.a.3", "-1", " 1", "1.3.6.1.2.1.4.22.1.3.1011.154.118.252.102",
	} {
		got, gotErr := ParseOID(in)
		want, wantErr := splitParseOID(in)
		if fmt.Sprint(gotErr) != fmt.Sprint(wantErr) {
			t.Errorf("%q: error %v, want %v", in, gotErr, wantErr)
			continue
		}
		if wantErr != nil {
			continue
		}
		if CompareOIDs(got, want) != 0 || len(got) != len(want) {
			t.Errorf("%q: %v, want %v", in, got, want)
		}
	}
}

func BenchmarkParseOID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustParseOID("1.3.6.1.2.1.4.35.1.4.1007.2.16.32.1.13.184.0.0.0.0.0.0.0.0.0.1.2.3")
	}
}
