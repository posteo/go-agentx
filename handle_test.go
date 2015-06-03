package agentx_test

import (
	"testing"

	"github.com/posteo/go-agentx/pdu"
	. "github.com/posteo/go-agentx/test"
	"github.com/posteo/go-agentx/value"
)

func TestGet(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	session.GetHandler = func(oid value.OID) (value.OID, pdu.VariableType, interface{}, error) {
		if oid.String() == "1.3.6.1.4.1.8072.3.1" {
			return oid, pdu.VariableTypeOctetString, "test", nil
		}
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}

	AssertNoError(t,
		session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

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
	session.GetNextHandler = func(from, to value.OID) (value.OID, pdu.VariableType, interface{}, error) {
		if from.String() == "1.3.6.1.4.1.8072.3.1" {
			return from, pdu.VariableTypeOctetString, "test", nil
		}
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}

	AssertNoError(t,
		session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

	AssertEquals(t,
		".1.3.6.1.4.1.8072.3.1 = STRING: \"test\"",
		SNMPGetNext(t, "1.3.6.1.4.1.8072.3.1"))
}

func TestGetBulk(t *testing.T) {
	session, err := e.client.Session()
	AssertNoError(t, err)
	defer session.Close()
	session.GetNextHandler = func(from, to value.OID) (value.OID, pdu.VariableType, interface{}, error) {
		if from.String() == "1.3.6.1.4.1.8072.3.1" {
			return from, pdu.VariableTypeOctetString, "test", nil
		}
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}

	AssertNoError(t,
		session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

	AssertEquals(t,
		".1.3.6.1.4.1.8072.3.1 = STRING: \"test\"",
		SNMPGetBulk(t, "1.3.6.1.4.1.8072.3.1", 0, 1))
}
