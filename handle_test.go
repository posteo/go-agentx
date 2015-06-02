package agentx_test

import (
	"testing"

	"github.com/posteo/go-agentx"
	"github.com/posteo/go-agentx/pdu"
)

func TestGet(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	session.GetHandler = func(oid string) (*agentx.Item, error) {
		if oid == "1.3.6.1.4.1.8072.3.1" {
			return &agentx.Item{Type: pdu.VariableTypeOctetString, Value: "test"}, nil
		}
		return nil, nil
	}

	AssertNoError(t,
		session.Register(127, "1.3.6.1.4.1.8072"))

	AssertEquals(t,
		".1.3.6.1.4.1.8072.3.1 = STRING: \"test\"",
		SNMPGet(t, "1.3.6.1.4.1.8072.3.1"))

	AssertEquals(t,
		".1.3.6.1.4.1.8072.3.2 = No Such Object available on this agent at this OID",
		SNMPGet(t, "1.3.6.1.4.1.8072.3.2"))
}

func TestGetNext(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	session.GetNextHandler = func(from, to string) (*agentx.Item, error) {
		if from == "1.3.6.1.4.1.8072.3.1" {
			return &agentx.Item{OID: "1.3.6.1.4.1.8072.3.1", Type: pdu.VariableTypeOctetString, Value: "test"}, nil
		}
		return nil, nil
	}

	AssertNoError(t,
		session.Register(127, "1.3.6.1.4.1.8072"))

	AssertEquals(t,
		".1.3.6.1.4.1.8072.3.1 = STRING: \"test\"",
		SNMPGetNext(t, "1.3.6.1.4.1.8072.3.1"))
}
