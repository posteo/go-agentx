package agentx_test

import (
	"testing"

	"github.com/juju/errgo"
)

func TestSessionOpen(t *testing.T) {
	session, err := e.client.Session()
	if err != nil {
		t.Fatalf(errgo.Details(err))
	}
	defer session.Close()

	if session.ID() == 0 {
		t.Fatalf("expected session id, got 0")
	}
}

func TestSessionClose(t *testing.T) {
	session, _ := e.client.Session()
	if err := session.Close(); err != nil {
		t.Fatalf(errgo.Details(err))
	}
}

func TestSessionRegister(t *testing.T) {
	session, _ := e.client.Session()
	defer session.Close()

	if err := session.Register(127, "1.3.6.1.4.1.8072"); err != nil {
		t.Fatalf(errgo.Details(err))
	}
}

func TestSessionAllocateIndex(t *testing.T) {
	session, _ := e.client.Session()
	defer session.Close()
	session.Register(127, "1.3.6.1.4.1.8072")

	if err := session.AllocateIndex("1.3.6.1.4.1.8072.1"); err != nil {
		t.Fatalf(errgo.Details(err))
	}
}
