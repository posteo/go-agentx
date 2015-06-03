package agentx_test

import (
	"testing"

	. "github.com/posteo/go-agentx/test"
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
