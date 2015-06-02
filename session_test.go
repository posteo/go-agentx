package agentx_test

import "testing"

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

func TestSessionRegister(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()

	err = session.Register(127, "1.3.6.1.4.1.8072")
	AssertNoError(t, err)
}

func TestSessionAllocateIndex(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	err = session.Register(127, "1.3.6.1.4.1.8072")
	AssertNoError(t, err)

	err = session.AllocateIndex("1.3.6.1.4.1.8072.1")
	AssertNoError(t, err)
}
