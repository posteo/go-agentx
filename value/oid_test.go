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
