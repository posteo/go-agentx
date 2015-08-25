/*
go-agentx
Copyright (C) 2015 Philipp Brüll <bruell@simia.tech>

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301
USA
*/

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
