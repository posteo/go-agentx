// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx_test

import (
	"testing"

	. "github.com/posteo/go-agentx"
	"github.com/posteo/go-agentx/pdu"
	. "github.com/posteo/go-agentx/test"
)

var listHandler = &ListHandler{}

func init() {
	item := listHandler.Add("1.3.6.1.4.1.45995.3.1")
	item.Type = pdu.VariableTypeOctetString
	item.Value = "test"
}

func TestGet(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	session.Handler = listHandler

	AssertNoError(t,
		session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

	AssertEquals(t,
		".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
		SNMPGet(t, "1.3.6.1.4.1.45995.3.1"))

	AssertEquals(t,
		".1.3.6.1.4.1.45995.3.2 = No Such Object available on this agent at this OID",
		SNMPGet(t, "1.3.6.1.4.1.45995.3.2"))
}

func TestGetNext(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	session.Handler = listHandler

	AssertNoError(t,
		session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

	AssertEquals(t,
		".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
		SNMPGetNext(t, "1.3.6.1.4.1.45995.3.1"))
}

func TestGetNextForChildOID(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	session.Handler = listHandler

	AssertNoError(t,
		session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

	AssertEquals(t,
		".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
		SNMPGetNext(t, "1.3.6.1.4.1.45995.3"))
}

func TestGetBulk(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	session.Handler = listHandler

	AssertNoError(t,
		session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

	AssertEquals(t,
		".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
		SNMPGetBulk(t, "1.3.6.1.4.1.45995.3.1", 0, 1))
}
