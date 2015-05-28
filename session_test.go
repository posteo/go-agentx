package agentx_test

import (
	"testing"

	"github.com/juju/errgo"
)

func TestSessionOpen(t *testing.T) {
	session, err := e.client.Session()
	if err != nil {
		t.Fatalf("error %s", errgo.Details(err))
	}
	defer session.Close()

	if session.ID() == 0 {
		t.Fatalf("expected session id, got 0")
	}
}

func TestSessionClose(t *testing.T) {
	session, _ := e.client.Session()
	if err := session.Close(); err != nil {
		t.Fatalf("error %s", errgo.Details(err))
	}
}
