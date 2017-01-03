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

	. "github.com/martinclaro/go-agentx/test"
)

func TestSessionOpen(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()

	AssertNotEquals(t, 0, session.ID())
}

func TestSessionClose(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)

	err = session.Close()
	AssertNoError(t, err)
}

func TestSessionRegistration(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()

	AssertNoError(t,
		session.Register(127, baseOID))

	AssertNoError(t,
		session.Unregister(127, baseOID))
}
