package agentx_test

import (
	"testing"

	"github.com/posteo/go-agentx"
	"github.com/posteo/go-agentx/pdu"
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
		session.Register(127, "1.3.6.1.4.1.8072"))

	AssertNoError(t,
		session.Unregister(127, "1.3.6.1.4.1.8072"))
}

func TestSessionIndexAllocation(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()

	AssertNoError(t,
		session.Register(127, "1.3.6.1.4.1.8072"))
	defer session.Unregister(127, "1.3.6.1.4.1.8072")

	AssertNoError(t,
		session.AllocateIndex(&agentx.Item{OID: "1.3.6.1.4.1.8072.3.1", Type: pdu.VariableTypeInteger, Value: int32(123)}))

	AssertNoError(t,
		session.DeallocateIndex(&agentx.Item{OID: "1.3.6.1.4.1.8072.3.1", Type: pdu.VariableTypeInteger, Value: int32(123)}))
}
